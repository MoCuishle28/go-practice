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
		for name in route.near_routes:
			table[0] = name
			table.append(route.name)
			s.send(';'.join(table).encode('utf-8'))
			table.pop()

		data = s.recv(4096).decode('utf-8').split(";")
		data.pop(0)
		print(data)

main()