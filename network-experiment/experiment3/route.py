import socket
import threading
import json


class Item(object):
	def __init__(self, target, distance, nextJump):
		self.target = target
		self.distance = distance
		self.nextJump = nextJump


	def __str__(self):
		return "("+str(self.target)+","+str(self.distance)+","+str(self.nextJump)+")"


class Route(object):
	def __init__(self, name):
		self.name = name
		self.network = set()		# 相邻的网络
		self.near_routes = set()	# 相邻的路由
		self.table = {}				# 路由表 (目的网络:Item)


def main():
	uid = input()
	s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
	s.connect(('127.0.0.1',9999))
	s.send(uid.encode('utf-8'))

	with open(uid+'.json', 'r') as f:
		data = json.load(f)
	
	route = Route(uid)
	for k,v in data.items():
		print(k, v)
		if k == 'network':
			route.network = set(v)
		elif k == 'near_routes':
			route.near_routes = set(v)
		elif k == 'table':
			for x in v:
				route.table[x[0]] = Item(x[0], x[1], x[2])

	print(route.network)
	print(route.near_routes)
	for k, v in route.table.items():
		print(k, v)

	change = True
	input()
	print("start...")
	while change:
		change = False
		table = [None]
		for _, item in route.table.items():
			tmp = [item.target, item.distance, item.nextJump]
			table.append(','.join(tmp))
		table.append(route.name)
		for name in route.near_routes:
			table[0] = name
			fragement = ';'.join(table)+"EOF"
			s.send(fragement.encode('utf-8'))

		# 可能同时受到多个route发来的表 以EOF为分割
		data = s.recv(4096).decode('utf-8').split("EOF")
		data.remove('')
		# print(data)

		# 修改路由表
		for form in data:
			form = form.split(";")
			form.pop(0)
			source_route = form.pop()
			for x in form:
				x = x.split(',')	# 0->目的网络, 1->距离, 2->下一跳
				x[-1] = source_route
				x[1] = str(int(x[1]) + 1)

				# 和当前路由表比对
				if x[0] not in route.table:		# 若不存在目的网络
					route.table[x[0]] = Item(x[0], x[1], x[2])
					change = True
				else:
					item = route.table[x[0]]
					if item.nextJump == x[2]:	# 若下一跳相同
						route.table[x[0]] = Item(x[0], x[1], x[2])
						change = True
					else:						# 若下一跳不同
						if int(item.distance) <= int(x[1]):
							route.table[x[0]] = item
						else:
							route.table[x[0]] = Item(x[0], x[1], x[2])
							change = True

		for _, item in route.table.items():
			print(item)
		print("----END----")

main()