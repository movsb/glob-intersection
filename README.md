# Glob-Intersection

This example code is to detect if two [glob][]-patterns have intersections.

Here we say that if two glob patterns intersect when they can match at least one path.
We do not aim to tell you that path.

[glob]: https://en.wikipedia.org/wiki/Glob_\(programming\)
[match]: https://pkg.go.dev/path/filepath@go1.18.3#Match

This is not like what [`filepath.Match`][match] does.
While `filepath.Match` tests glob against path, we glob-intersection tests glob against glob.

## Caveats

* Single path component only
* Only `*` is supported
* `**` is not supported
* `?` is not supported
* `[]` is not supported
* `\` is not supported

## Tests

Below is an example of trying to detect if `ab*d` and `a*cd` can match some path.

```-
|           ab*d| |           a*cd| initial
|            b*d| |            *cd| advance both: a ↔ a
|            b*d| |             cd| expand right: * → ∅
|             *d| |             cd| expand right: * → b
|              d| |             cd| expand  left: * → ∅
|              d| |              d| expand  left: * → c
|               | |               | advance both: d ↔ d
|               | |               | intersected
```

Each line is an action with the states of the rune stack after that action done.
The first line is the initial states. And at the last line, both stack go empty, we know those glob-patterns are intersected.

## References

* [regex - Algorithm to find out whether the matches for two Glob patterns (or Regular Expressions) intersect - Stack Overflow](https://stackoverflow.com/q/18695727/3628322)
* [Pathgather/glob-intersection: Find the intersection of two glob patterns](https://github.com/Pathgather/glob-intersection)

## More Tests

```shell-session
$ go test -run ^TestIntersection$ github.com/movsb/glob-intersection
----------------------------------------------------------------------
|             ab| |              a| initial
|              b| |               | advance both: a ↔ a
|             ab| |              a| not intersected
----------------------------------------------------------------------
|              *| |          *.log| initial
|               | |          *.log| expand  left: * → ∅
|               | |           .log| expand right: * → ∅
|               | |           .log| expand  left: * → *
|               | |            log| expand  left: * → *.
|               | |             og| expand  left: * → *.l
|               | |              g| expand  left: * → *.lo
|               | |               | expand  left: * → *.log
|              *| |          *.log| intersected
----------------------------------------------------------------------
|              *| |            abc| initial
|               | |            abc| expand  left: * → ∅
|               | |             bc| expand  left: * → a
|               | |              c| expand  left: * → ab
|               | |               | expand  left: * → abc
|              *| |            abc| intersected
----------------------------------------------------------------------
|          *.log| |          a.log| initial
|           .log| |          a.log| expand  left: * → ∅
|           .log| |           .log| expand  left: * → a
|            log| |            log| advance both: . ↔ .
|             og| |             og| advance both: l ↔ l
|              g| |              g| advance both: o ↔ o
|               | |               | advance both: g ↔ g
|          *.log| |          a.log| intersected
----------------------------------------------------------------------
|       *-live-*| |         -live-| initial
|        -live-*| |         -live-| expand  left: * → ∅
|         live-*| |          live-| advance both: - ↔ -
|          ive-*| |           ive-| advance both: l ↔ l
|           ve-*| |            ve-| advance both: i ↔ i
|            e-*| |             e-| advance both: v ↔ v
|             -*| |              -| advance both: e ↔ e
|              *| |               | advance both: - ↔ -
|               | |               | expand  left: * → ∅
|       *-live-*| |         -live-| intersected
----------------------------------------------------------------------
|       *-live-*| |        a-live-| initial
|        -live-*| |        a-live-| expand  left: * → ∅
|        -live-*| |         -live-| expand  left: * → a
|         live-*| |          live-| advance both: - ↔ -
|          ive-*| |           ive-| advance both: l ↔ l
|           ve-*| |            ve-| advance both: i ↔ i
|            e-*| |             e-| advance both: v ↔ v
|             -*| |              -| advance both: e ↔ e
|              *| |               | advance both: - ↔ -
|               | |               | expand  left: * → ∅
|       *-live-*| |        a-live-| intersected
----------------------------------------------------------------------
|       *-live-*| |        -live-b| initial
|        -live-*| |        -live-b| expand  left: * → ∅
|         live-*| |         live-b| advance both: - ↔ -
|          ive-*| |          ive-b| advance both: l ↔ l
|           ve-*| |           ve-b| advance both: i ↔ i
|            e-*| |            e-b| advance both: v ↔ v
|             -*| |             -b| advance both: e ↔ e
|              *| |              b| advance both: - ↔ -
|               | |              b| expand  left: * → ∅
|               | |               | expand  left: * → b
|       *-live-*| |        -live-b| intersected
----------------------------------------------------------------------
|       *-live-*| |       a-live-b| initial
|        -live-*| |       a-live-b| expand  left: * → ∅
|        -live-*| |        -live-b| expand  left: * → a
|         live-*| |         live-b| advance both: - ↔ -
|          ive-*| |          ive-b| advance both: l ↔ l
|           ve-*| |           ve-b| advance both: i ↔ i
|            e-*| |            e-b| advance both: v ↔ v
|             -*| |             -b| advance both: e ↔ e
|              *| |              b| advance both: - ↔ -
|               | |              b| expand  left: * → ∅
|               | |               | expand  left: * → b
|       *-live-*| |       a-live-b| intersected
----------------------------------------------------------------------
|           ab*d| |           a*cd| initial
|            b*d| |            *cd| advance both: a ↔ a
|            b*d| |             cd| expand right: * → ∅
|             *d| |             cd| expand right: * → b
|              d| |             cd| expand  left: * → ∅
|              d| |              d| expand  left: * → c
|               | |               | advance both: d ↔ d
|           ab*d| |           a*cd| intersected
----------------------------------------------------------------------
|        abc*def| |        abc.def| initial
|         bc*def| |         bc.def| advance both: a ↔ a
|          c*def| |          c.def| advance both: b ↔ b
|           *def| |           .def| advance both: c ↔ c
|            def| |           .def| expand  left: * → ∅
|            def| |            def| expand  left: * → .
|             ef| |             ef| advance both: d ↔ d
|              f| |              f| advance both: e ↔ e
|               | |               | advance both: f ↔ f
|        abc*def| |        abc.def| intersected
----------------------------------------------------------------------
|         abcd*g| |         a*defg| initial
|          bcd*g| |          *defg| advance both: a ↔ a
|          bcd*g| |           defg| expand right: * → ∅
|           cd*g| |           defg| expand right: * → b
|            d*g| |           defg| expand right: * → bc
|             *g| |            efg| advance both: d ↔ d
|              g| |            efg| expand  left: * → ∅
|              g| |             fg| expand  left: * → e
|              g| |              g| expand  left: * → ef
|               | |               | advance both: g ↔ g
|         abcd*g| |         a*defg| intersected
----------------------------------------------------------------------
|            *bc| |            ab*| initial
|             bc| |            ab*| expand  left: * → ∅
|             bc| |             b*| expand  left: * → a
|              c| |              *| advance both: b ↔ b
|              c| |               | expand right: * → ∅
|               | |               | expand right: * → c
|            *bc| |            ab*| intersected
----------------------------------------------------------------------
|              *| |             **| initial
|               | |             **| expand  left: * → ∅
|               | |              *| expand right: * → ∅
|               | |               | expand right: * → ∅
|              *| |             **| intersected
----------------------------------------------------------------------
|        a*b*c*d| |             **| initial
|        a*b*c*d| |              *| expand right: * → ∅
|        a*b*c*d| |               | expand right: * → ∅
|         *b*c*d| |               | expand right: * → a
|          b*c*d| |               | expand  left: * → ∅
|          b*c*d| |               | expand right: * → a*
|           *c*d| |               | expand right: * → a*b
|            c*d| |               | expand  left: * → ∅
|            c*d| |               | expand right: * → a*b*
|             *d| |               | expand right: * → a*b*c
|              d| |               | expand  left: * → ∅
|              d| |               | expand right: * → a*b*c*
|               | |               | expand right: * → a*b*c*d
|        a*b*c*d| |             **| intersected
----------------------------------------------------------------------
|        a*b*c*d| |      *********| initial
|        a*b*c*d| |       ********| expand right: * → ∅
|        a*b*c*d| |        *******| expand right: * → ∅
|        a*b*c*d| |         ******| expand right: * → ∅
|        a*b*c*d| |          *****| expand right: * → ∅
|        a*b*c*d| |           ****| expand right: * → ∅
|        a*b*c*d| |            ***| expand right: * → ∅
|        a*b*c*d| |             **| expand right: * → ∅
|        a*b*c*d| |              *| expand right: * → ∅
|        a*b*c*d| |               | expand right: * → ∅
|         *b*c*d| |               | expand right: * → a
|          b*c*d| |               | expand  left: * → ∅
|          b*c*d| |               | expand right: * → a*
|           *c*d| |               | expand right: * → a*b
|            c*d| |               | expand  left: * → ∅
|            c*d| |               | expand right: * → a*b*
|             *d| |               | expand right: * → a*b*c
|              d| |               | expand  left: * → ∅
|              d| |               | expand right: * → a*b*c*
|               | |               | expand right: * → a*b*c*d
|        a*b*c*d| |      *********| intersected
--- PASS: TestIntersection (0.00s)
PASS
ok  	github.com/movsb/glob-intersection	0.660s
```
