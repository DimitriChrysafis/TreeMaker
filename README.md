# my first go project
## what this is
This is a playground which makes trees of your desired depth.
  - you are able to add nodes to the graph event after it is generated
  - it makes a graph.txt file if you want to hang onto it after you make it
Sample graph (of depth input 3):
```
                         |-- 12
                 |-- 8
                         |-- 11
         |-- 3
                         |-- 10
                 |-- 7
                         |-- 9
 |-ROOT 1
                         |-- 6
                 |-- 2
                         |-- 5
         |-- 2
                         |-- 4
                 |-- 1
                         |-- 3

```

after the graph is outputted

```
Do you want to add a child node? (yes/no):
yes
Which side do you want to add the child to? (left/right):
right
Enter the data of the target node:
5
Added node 13 to the right of node 5
Graph text file saved as graph.txt
Your tree:
                         |-- 12
                 |-- 8
                         |-- 11
         |-- 3
                         |-- 10
                 |-- 7
                         |-- 9
 |-- 1
                         |-- 6
                 |-- 2
                                 |-- 13
                         |-- 5
         |-- 2
                         |-- 4
                 |-- 1
                         |-- 3

```
