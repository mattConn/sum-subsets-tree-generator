#!/usr/bin/python

# This script receives a flat array of breadth-first binary tree data

import sys
import math
import networkx as nx
import matplotlib.pyplot as plt
import pygraphviz
from networkx.drawing.nx_agraph import graphviz_layout

# graph drawing options
def draw(graph):
    pos=graphviz_layout(graph, prog='dot')
    nx.draw_networkx(graph,pos,font_color='black',node_color='orange',font_size=15,node_size=700,)
    plt.title("depth:node:sum")

sys.argv[1:] = [int(i) for i in sys.argv[1:]]

depth = int(math.log(len(sys.argv[1:])+1,2))
maxNodesAtDepth = 2**(depth-1)

# plotting options
plt.figure(figsize=(maxNodesAtDepth,depth))

# construct graph
g = nx.DiGraph();

#for i in sys.argv[1:]:
#    g.add_node(i)

for dataIndex in range(len(sys.argv[1:])):
    leftIndex = 2*dataIndex+1
    rightIndex = 2*dataIndex+2

    if leftIndex >= len(sys.argv[1:]): break 

    def getDepth(index): return int(math.log(index+1,2))

    # for connecting current node to another
    # labelling each node as "dataIndex:data[dataIndex]"
    def joinCurNodeTo(otherIndex):
        g.add_edge(f'{getDepth(dataIndex)}:{dataIndex}:{sys.argv[1:][dataIndex]}', f'{getDepth(otherIndex)}:{otherIndex}:{sys.argv[1:][otherIndex]}')

    joinCurNodeTo(leftIndex)
    joinCurNodeTo(rightIndex)
    

draw(g)
plt.savefig("assets/graph.png")
