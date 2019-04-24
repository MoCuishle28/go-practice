import socket
import threading, time


def main():
	s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
	s.connect(('127.0.0.1',9999))
	msg = input()
	s.send(msg.encode('utf-8'))

	for i in range(3):
		msg = input()
		s.send(msg.encode('utf-8'))
		print(s.recv(4096).decode('utf-8'))

	s.close()


main()