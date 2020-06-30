#!/usr/bin/python

# This script receives a flat array of breadth-first binary tree data

import sys
import math
import networkx as nx
import matplotlib.pyplot as plt
import pygraphviz
from networkx.drawing.nx_agraph import graphviz_layout

# graph drawing
def draw(graph):
    pos=graphviz_layout(graph, prog='dot')
    nx.draw_networkx(graph,pos,font_color='black',node_color='orange',font_size=15,node_size=700,)
    plt.title("node:sum")

# cast argv (binary tree data) to int
treeData = [int(i) for i in sys.argv[1:]]

# dimensions of tree
depth = int(math.log(len(treeData)+1,2))
maxNodesAtDepth = 2**(depth-1)

# size plot by tree dimensions
plt.figure(figsize=(maxNodesAtDepth,depth))

# construct graph
g = nx.DiGraph();

# traverse tree data array and add edges to networkx graph
for dataIndex in range(len(treeData)):
    leftIndex = 2*dataIndex+1 # index of left node in tree
    rightIndex = 2*dataIndex+2 # index of right node in tree

    if leftIndex >= len(treeData): break # break if end of data

    # for connecting current node to another node.
    # labelling each node as "dataIndex:data[dataIndex]"
    def joinCurrentNodeTo(otherIndex):
        g.add_edge(f'{dataIndex}:{treeData[dataIndex]}', f'{otherIndex}:{treeData[otherIndex]}')

    joinCurrentNodeTo(leftIndex)
    joinCurrentNodeTo(rightIndex)
    

draw(g) # draw graph
plt.savefig("assets/graph.png") # save graph image
