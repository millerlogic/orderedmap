# orderedmap
Go ordered map.
Documentation: [godoc](https://godoc.org/github.com/millerlogic/orderedmap)

Can be used with [go code generator](https://github.com/reusee/ccg) to generate a type safe ordered map, such as:
```
ccg -f github.com/millerlogic/orderedmap -t KeyType=int,ValueType=string -r OrderedMap=IntStringOrderedMap,NewOrderedMap=NewIntStringOrderedMap,link=intstringlink -o intstringorderedmap.go
```
