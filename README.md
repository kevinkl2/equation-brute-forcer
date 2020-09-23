# equation-brute-forcer

Bogosort implementation of equation brute forcer.

Given 9 variables, their values, and a final result, produce an equation that fits those parameters.

For example:
```python
values = {"a":"10", "b":"5", "c":"3", "d":"1.5", "e":"1.75", "f":"15", "g":"20", "h":"1.02", "i":"2"}
final result = 85.895
equation = h*(d*e*(a + b + c) + f + g) + i
```