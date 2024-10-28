# Experiment different dependency injection frameworks

While Dependency Injection and IoC is first-class paradigm in GO, many "Gophers" paradoxically dislike DI frameworks and DI containers. Their arguments are usually:

* you don't need a DI framework because Go itself provided enough support to practice DI.
* these frameworks are too "magical", go against Go idiomatique.
* reflection are bad, hurts perf.

In my case, I have to work with complex application with a complex dependencies such as:

```mermaid
flowchart TD
A["A<br><sup>(transient)</sup>"]
B["B<br><sup>(transient)</sup>"]
C["C<br><sup>(scoped)</sup>"]
D["D<br><sup>(transient)<br>(shutdowner)</sup>"]
E["E<br><sup>(scoped)<br>(shutdowner)</sup>"]
F["F<br><sup>(transient)</sup>"]
G(["G<br><sup>(interface)</sup>"])
Gb("Gb<br><sup>(scoped)</sup>")
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

Manually write and maintain codes like this `a := NewA(NewB(NewD(NewF(), NewH()), NewE([]G{NewDGa(NewGa()), NewGb(), NewGc()})), NewC())` to wire things together is a "No, thanks.." for me.

=> I need helps from a DI framework:

* to wire thing together and to provide the right `A`, `B`, `C`.. object for me whenever I need them (magically or not I don't care).
* to minimize the wiring work when `A` or `D` get more dependency in the future...

Lastly, spending some more nano-second on Reflection is the not the things I would have to worry about.

With this use case in mind, I tried some popular DI frameworks to select the right one for me.

<https://uber-go.github.io/fx/get-started/>

* (+) non-instrusive drop-in: use or remove the library on exising codes codes requires zero or minimal refactor / changes
* (+) very popular
* (+) support multi-implementation, decorator, module..
* (-) singleton only, transient injection not possible

<https://github.com/golobby/container/>

* (-) do not support multi-implementation (the last registered implementation takes precedent)
* (-) instrusive integration / tightly coupled library: use or remove the library on exising codes requires refactor/changes
  * binding requires adding specialize constructor which return the Abstraction
  * auto-wiring (aka `Fill` requires adding `container` tag to the existing codes)
* (-) nobody answer question on github repo

<https://github.com/samber/do>

* (+) benefit from Go generics
* (+) non instrusive integration is possible (the provider will be quite verbose in this case)
  * otherwise: instrusive integration / tightly coupled library: adding specialized constructor which take "do.injector" as input
  * automatic wiring requires adding tag `do:""` to the struct (a little slower because of reflection)
* (+) support transient
* (+) support package, module
* (+) possible to use module for scoped life time
* (-) multi-implementation injection is not possible (a random matching will be choosen)
  * (+) [this feature is on the way to V2](https://github.com/samber/do/pull/45)
* (-) development seem not very active
* (+) shutdowner, healthcheck

<https://github.com/firasdarwish/ore>

* (+) benefit from Go generics
* (+) non instrusive integration is possible
  * Registering codes are very verbose and explicit, not much different from samber/do
  * (-) no automatic wiring => (+) no reflection
* (+) support transient & Scoped via context => a very natural choice for Go
* (+) multiple implementation injection is possible
* (+) very light-weight: more performance, use less memory than samber/do
* (-) no handy shutdowner, healthchecker as samber/do
* (-) too young, not popular
* (-) [unable to get ore.scoped works on the complex sample](https://github.com/firasdarwish/ore/issues/2).

=> my choice: samber/do because

* I will need "transient" object injection which uber/fx doesn't have.
* Golobby and other library have great potential but are too unpopular to bet on.
* I will need the handy shutdowner.. with samber/do, I don't have to handle it by myself

## Run benchmark

```sh
go test -benchmem -bench . try-uberfx/demo1
```
