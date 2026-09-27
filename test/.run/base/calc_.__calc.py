#!/usr/bin/env python3
# A python test script: reports args and runs a trivial computation.
import sys
print("PYTHON: argc =", len(sys.argv) - 1)
for i, a in enumerate(sys.argv[1:]):
    print(f"PYTHON: arg[{i}] = {a}")
print("PYTHON: sum =", sum(range(10)))
