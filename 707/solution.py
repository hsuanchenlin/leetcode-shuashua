class Node(object):

    def __init__(self, val):
        self.val = val
        self.next = None
        
class MyLinkedList:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.head = None
        self.tail = None
        self.size = 0

    def get(self, index: int) -> int:
        """
        Get the value of the index-th node in the linked list. If the index is invalid, return -1.
        """
        if index < 0 or index >= self.size:
            return -1
        node = self.head
  
        if not node:
            return -1  
  
        for i in range(index):
            # print("index ",i,self.size)
            node = node.next
        
        return node.val

    def addAtHead(self, val: int) -> None:
        """
        Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
        """
        node = Node(val)
        if not self.head:
            self.tail = node
            self.head = node
        else:
            node.next = self.head
            self.head = node
        self.size += 1
        
    
    def addAtTail(self, val: int) -> None:
        """
        Append a node of value val to the last element of the linked list.
        """
        node = Node(val)
        if not self.tail:
            self.tail = node
            self.head = node
        else:
            self.tail.next = node
            self.tail = node
        self.size += 1
        
    def addAtIndex(self, index: int, val: int) -> None:
        """
        Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
        """
        if index < 0 or index > self.size:
            return -1
        if index == 0:
            self.addAtHead(val)
            return
        if index == self.size:
            self.addAtTail(val)
            return 
        if index != 0 and not self.head:
            return
      
        nn = Node(val)
       
        prev = self.head
        for i in range(index-1):   
            prev = prev.next
        nn.next = prev.next
        prev.next = nn
        self.size +=1
        
    def deleteAtIndex(self, index: int) -> None:
        """
        Delete the index-th node in the linked list, if the index is valid.
        """
        if index >= self.size or index < 0:
            return 
        if not self.head:
            return
       
        node = self.head
        prev = None
        if index == 0:
            self.head = self.head.next
            self.size -= 1
            return
        
        for i in range(index-1):  
            node = node.next
        if index == self.size-1:
            node.next = None
            self.tail = node
        else:
            node.next = node.next.next
        self.size -= 1
        

# Your MyLinkedList object will be instantiated and called as such:
# obj = MyLinkedList()
# param_1 = obj.get(index)
# obj.addAtHead(val)
# obj.addAtTail(val)
# obj.addAtIndex(index,val)
# obj.deleteAtIndex(index)