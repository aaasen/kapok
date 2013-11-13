
# Kapok

A Knowledge Graph of Wikipedia.

![](https://raw.github.com/aaasen/kapok/master/demo/anarchism.png)

## Description

Kapok aims to create a knowledge graph from Wikipedia.
In this graph, each node is an article, and links between articles are the
edges between nodes.

## Structure

Kapok is split into 3 modular sections:

 - **Parsing:** extracting relevant data from a 45GB archive of Wikipedia
 - **Graph:** morphing the parsed data into a graph for analysis
 - **Visualisation:** creating interesting visualisations with the data

The parsing section of Kapok could be easily extended to replace aging
Wikimedia tools like MWDumper. I'll probably do this soon.
