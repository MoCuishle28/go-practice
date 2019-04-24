import socket
import threading


def Transmit(sock, addr, conn_dict, uid):
	print('Accept new connection from %s...' % uid)
	while True:
		data = sock.recv(4096)
		if not data or data.decode('utf-8') == 'exit':
			print("exit:", not data)
			break

		recv_data = data.decode('utf-8')
		print("recv: ", recv_data)
		target_id = recv_data.split(';')[0]

		tmp_sock = conn_dict.get(target_id, [None])[0]
		if not tmp_sock:
			print(target_id, "is not exist!")
		else:
			tmp_sock.send(data)
	sock.close()
	del conn_dict[uid]
	print('%s closed.' % uid)


def main():
	s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
	s.bind(('127.0.0.1',9999))
	s.listen(10)	# 最大连接数 10
	print("Waiting for connection...")

	conn_dict = {}

	route_id = 0
	for i in range(10):
		# 接受一个新连接:
		sock, addr = s.accept()

		data = sock.recv(4096)	# 接受id
		route_id = data.decode('utf-8')	
		conn_dict[route_id] = (sock, addr)

		handle = threading.Thread(target=Transmit, args=(sock, addr, conn_dict, route_id) )
		handle.start()


main()