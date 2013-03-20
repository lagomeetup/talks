
import thread

def loopme(*args):
    while True:
        pass

def main():
    for i in xrange(1000):
        thread.start_new_thread(loopme, (None,))
    input()
    
main()

