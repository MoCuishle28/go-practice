import time


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
		self.table = {}		# 目的网络:Item


class Network(object):
	"""用 图(邻接矩阵) 来存储网络"""
	def __init__(self):
		"""节点是路由器, 边是网络"""
		self.edge = None
		self.routes = None	# 网络上的路由
		self.init_network()


	def init_network(self):
		self.routes = {i:Route(chr(i + 65)) for i in range(6)}		# 0~5 代表 A~F
		self.edge = [
		#	A   B  C  D  E  F
			[0, 3, 0, 2, 1, 0],		# A
			[3, 0, 4, 0, 0, 0],		# B	
			[0, 4, 0, 0, 0, 6],		# C
			[2, 0, 0, 0, 5, 5],		# D
			[1, 0, 0, 5, 0, 5],		# E
			[0, 0, 6, 5, 5, 0]		# F
		]

		for k, v in self.routes.items():
			for i, net in enumerate(self.edge[k]):
				if net == 0:
					continue
				v.table[net] = Item(net, 1, '_')


	def RIP(self):
		pass


def main():
	network = Network()
	print("init...")
	for name, route in network.routes.items():
		print(route.name, ":")
		for k,v in route.table.items():
			print(k, v)
		print('------')

	print("开始模拟RIP...")
	network.RIP()
	print("模拟结束...")

	# TODO 模拟坏消息


main()