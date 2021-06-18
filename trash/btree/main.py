import json

class Node:
    def __init__(self, data):

        self.left = None
        self.right = None
        self.data = data

    def PrintTree(self):
        if self.left:
            self.left.PrintTree()
        print( self.data),
        if self.right:
            self.right.PrintTree()

    def insert(self, data):
# Compare the new value with the parent node
        if self.data:
            if data < self.data:
                if self.left is None:
                    self.left = Node(data)
                else:
                    self.left.insert(data)
            elif data > self.data:
                if self.right is None:
                    self.right = Node(data)
                else:
                    self.right.insert(data)
        else:
            self.data = data
            
class Node:

    def __init__(self,data):
        self.values = []
        self.keys = []
        self.size = 5

        self.leaf = True

    def split(self):
        

    def insert(self,value):
        if self.leaf == True:
            if len(self.values) =< self.size:
                self.values.append(value)
                self.values.sort()
            else:

        else:
            for x in values:

