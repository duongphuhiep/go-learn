# Dependency graph

```mermaid
flowchart TD
A["A<br><sup>(transient)</sup>"]
B["B<br><sup>(transient)</sup>"]
D["D<br><sup>(transient)<br>(shutdowner)</sup>"]
E["E<br><sup>(shutdowner)</sup>"]
F["F<br><sup>(transient)</sup>"]
G(["G<br><sup>(interface)</sup>"])
Ga("Ga<br><sup>(shutdowner)</sup>")
DGa("DGa<br><sup>(decorator)</sup>")
H["H<br><sup>(shutdowner)</sup>"]

A-->B
A-->C
B-->D
B-->E
D-->F
D-->H
E-->DGa
E-->Gb
E-->Gc
DGa-->|decorate| Ga
Ga -.implement..-> G
Gb -.implement..-> G
Gc -.implement..-> G
DGa -.implement..-> G
```
