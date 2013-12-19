
Create two nodes and relate them. Won't make extra nodes.

```
MERGE (a:Article { title:"a" })  MERGE (b:Article { title:"b" }) 
CREATE UNIQUE (a)-[r:REFERS_TO]->(b) 
RETURN a, r, b
```

Create a and b if they do not exist and add categories to a.

```
MERGE (a:Article { title:"a" })
MERGE (b:Article { title:"b" })
FOREACH (catName IN ["cat1", "cat2"]| MERGE (cat:Category { title:catName }) 
	CREATE UNIQUE (a)-[:IN_CATEGORY]-(cat)) 
CREATE UNIQUE (a)-[r:REFERS_TO]->(b) 
RETURN a
```

Create an origin article and add its neighbors and categories.

```
MERGE (origin:Article { title:"a" })
FOREACH (otherName IN ["b", "c", "d"] |
  MERGE (other:Article { title: otherName})
  CREATE UNIQUE (origin)-[r:REFERS_TO]->(other))
FOREACH (catName IN ["cat1", "cat2"] |
  MERGE (cat:Category { title:catName }) 
  CREATE UNIQUE (origin)-[:IN_CATEGORY]-(cat)) 
RETURN origin
```
