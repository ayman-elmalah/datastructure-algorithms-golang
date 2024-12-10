# Data structure and algorithms in [GOLANG](https://go.dev/) By [Ayman Elmalah](http://www.ayman-elmalah.com/)

Implement some data structure and algorithms topics in golang 

## You want to try ?

- Clone repo
  `` git clone https://github.com/ayman-elmalah/datastructure-algorithms-golang && cd datastructure-algorithms-golang``
- ``go mod tidy``
- Now you have directories in root directory, just run `` go run ${DIRECTORY_VAR}/main.go ``

## Content
- [Reverse a string](#reverse-a-string)
- [Which repeated first](#which-repeated-first)
- [Merge sorted arrays](#merge-sorted-arrays)
- [Merge sorted lists](#merge-sorted-lists)
- [Linked list](#linked-list)
- [Stack](#stack)
- [Queue](#queue)
- [Binary tree](#binary-tree)
- [Graph](#graph)
- [Factorial](#factorial)
- [Fibonacci](#fibonacci)
- [Bubble sort](#bubble-sort)
- [Selection sort](#selection-sort)
- [Insertion sort](#insertion-sort)
- [Merge sort](#merge-sort)
- [Linear search](#linear-search)
- [Binary search](#binary-search)
- [Longest common prefix](#longest-common-prefix)
- [Digital clock](#digital-clock)
- [Find the Missing Number in a Sequence](#find-the-missing-number-in-a-sequence)

<hr>

### Reverse a string
###### Question: Implement a solution to reverse given string characters, As example Given ``Hello world`` and the method should return ``dlrow olleH``

We handled it using 2 solutions
  - The first solution is to create a new variable and loop through the given string from the end then push the characters one by one to the new variable
  - The second solution is to loop through the middle of the given string then swap the first character with the last one, and the second with the previous of last character and so

You can have a look to code by viewing the file ``${root}/reverse_string/main.go``

Also you can run the program from the console ``go run reverse_string/main.go``

<hr>

### Which repeated first
###### Question: Implement a solution to detect which character repeated first, As example Given ``[1, 2, 3, 3, 5]`` so method should return ``3`` as it's the first character repeated in the array

We handled it using 2 solutions
- The first solution is to loop through the first array, then do a nested loop in the second so you can check which one repeated first
- The second solution is to use hash table data structure so you can just have one loop Big O notation will be O(n)

You can have a look to code by viewing the file ``${root}/which_repeated_first/main.go``

Also you can run the program from the console ``go run which_repeated_first/main.go``

<hr>

### Merge sorted arrays
###### Question: Implement a solution to merge 2 sorted arrays in 1 array, As example Given ``[0, 2, 4] and [1, 5, 6]`` so method should return ``[0, 1, 2, 4, 5, 6]`` as the sorted array

We handled it by comparing the first values in the 2 given arrays and push the smaller value to a new array then remove it from it's array and repeat it until the given arrays be empty

You can have a look to code by viewing the file ``${root}/merge_sorted_arrays/main.go``

Also you can run the program from the console ``go run merge_sorted_arrays/main.go``

<hr>

### Merge sorted lists
###### Question: Implement a solution to merge 2 sorted lists in 1 list, As example Given ``[0, 2, 4] and [1, 5, 6]`` so method should return ``[0, 1, 2, 4, 5, 6]`` as the sorted list

We handled it by comparing the first values in the 2 given lists and push the smaller value to a new list and repeat it until the given list is empty, the second solution used recursive

You can have a look to code by viewing the file ``${root}/merge_sorted_lists/main.go``

Also you can run the program from the console ``go run merge_sorted_lists/main.go``

<hr>

### Linked list
##### List that is linked, contains a set of nodes, each node have value and pointer
###### Implement linked list methods in golang methods: [Append, Prepend, Insert, Remove and Display]

You can have a look to code by viewing the file ``${root}/linked_list/main.go``

Also you can run the program from the console ``go run linked_list/main.go``

<hr>

### Stack
##### It's a linear data structure in which elements can be inserted or deleted from one side of list that is called TOP, It follows "LIFO" last in first out
###### Implement stack methods in golang methods: [Push, Pop, Length, Peek, and Display]

You can have a look to code by viewing the file ``${root}/stack/main.go``

Also you can run the program from the console ``go run stack/main.go``

<hr>

### Queue
##### It's a linear data structure in which elements can be inserted from one side that's called REAR and deleted from the other side that's called FRONT, It follows "FIFO" first in first out
###### Implement queue methods in golang methods: [Enqueue, Dequeue, Length, Peek and Display]

You can have a look to code by viewing the file ``${root}/queue/main.go``

Also you can run the program from the console ``go run queue/main.go``

<hr>

### Binary tree
#####  It's a very basic tree, each node can only have either 0,1 or 2 nodes and each child can have just 1 parent
###### Implement a binary tree methods in golang methods: [Insert and Lookup]

You can have a look to code by viewing the file ``${root}/binary_tree/main.go``

Also you can run the program from the console ``go run binary_tree/main.go``

<hr>

### Graph
#####  A graph data structure is a collection of nodes that have data and connected to other nodes, Nodes connected together with edges
###### Implement a graph methods in golang methods: [AddNode and AddEdge]

You can have a look to code by viewing the file ``${root}/graph/main.go``

Also you can run the program from the console ``go run graph/main.go``

<hr>

### Factorial
###### Implement a solution to find a factorial of number in golang using recursive and iterative

You can have a look to code by viewing the file ``${root}/factorial/main.go``

Also you can run the program from the console ``go run factorial/main.go``

<hr>

### Fibonacci
###### Implement a solution to find the number in the given index from fibonacci pattern

You can have a look to code by viewing the file ``${root}/fibonacci/main.go``

Also you can run the program from the console ``go run fibonacci/main.go``

<hr>

### Bubble sort
##### Bubble sort is the simplest sorting algorithm, it works by repeatedly swapping the adjacent elements if they are in the wrong order, biggest elements will be at the end   
###### Implement bubble sort algorithm to sort list of numbers

You can have a look to code by viewing the file ``${root}/bubble_sort/main.go``

Also you can run the program from the console ``go run bubble_sort/main.go``

<hr>

### Selection sort
##### Selection sort is a sorting algorithm that is working by repeatedly finding the smallest element in the array and put at the beginning
###### Implement selection sort algorithm to sort list of numbers

You can have a look to code by viewing the file ``${root}/selection_sort/main.go``

Also you can run the program from the console ``go run selection_sort/main.go``

<hr>

### Insertion sort
##### Insertion sort is a simple sorting algorithm that is working by finding element then put it in the right place, it can be efficient in small list  
###### Implement insertion sort algorithm to sort list of numbers

You can have a look to code by viewing the file ``${root}/insertion_sort/main.go``

Also you can run the program from the console ``go run insertion_sort/main.go``

<hr>

### Merge sort
##### Merge sort is a recursive sorting algorithm it's quite a bit faster that's dividing and conquer, divide input slides into 2 halves, recursively sort the two halves then merge sorted array 
###### Implement merge sort algorithm to sort list of numbers

You can have a look to code by viewing the file ``${root}/merge_sort/main.go``

Also you can run the program from the console ``go run merge_sort/main.go``

<hr>

### Linear search
##### A linear search is also known as a sequential search that simply scans each element at a time.  Suppose we want to search an element in an array or list, we simply calculate its length and do not jump at any item.
###### Implement linear search algorithm to search item in list

You can have a look to code by viewing the file ``${root}/linear_search/main.go``

Also you can run the program from the console ``go run linear_search/main.go``

<hr>

### Binary search
##### A binary search is a search in which the middle element is calculated It need to be sorted to check whether it is smaller or larger than the element which is to be searched. The main advantage of using binary search is that it does not scan each element in the list. Instead of scanning each element, it performs the searching to the half of the list. So, the binary search takes less time to search an element as compared to a linear search.
###### Implement binary search algorithm to search item in sorted list

You can have a look to code by viewing the file ``${root}/binary_search/main.go``

Also you can run the program from the console ``go run binary_search/main.go``

<hr>

### Longest common prefix
###### Question: Write a function to find the longest common prefix string amongst an array of strings, As example Given ``["Flowers", "Flows", "Flight"]`` and the method should return ``Fl``

You can have a look to code by viewing the file ``${root}/longest_prefix/main.go``

Also you can run the program from the console ``go run longest_prefix/main.go``

<hr>

### Digital clock
###### Question: Write a function to display a digital clock on terminal

You can have a look to code by viewing the file ``${root}/digital_clock/main.go``

Also you can run the program from the console ``go run digital_clock/main.go``

<hr>

### Find the Missing Number in a Sequence
###### Question: Write a function to find missing number from sequence array

You can have a look to code by viewing the file ``${root}/find_missing_number/main.go``

Also you can run the program from the console ``go run find_missing_number/main.go``