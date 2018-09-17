The way I see this code executing is by iterating over every character or group of letters
and calling the same funtion recursively and passing string minus the characters we iterated over
If we find a `[` opening bracket then we'd expect children for the current root, 
if we find `]` closing bracket or a `,` then this iteration has finished it's job and should return
if it's none of these, then this character which would be a letter is a sibling to the existing child
this root has.
We keep iterating until there is no longer any more characters left in the string.
That however, is an incomplete solution, because unlike most recursive functions, the nested functions
modify data that the parent functions need to continue their work, take the below example of how the flow
of data would go:

```
[a[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]
RECURSE
	a[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]
	[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]] <-------- *1
	RECURSE
		aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]
		[aaa],ab,ac],b,c[ca,cb,cc[cca]]]
		RECURSE
			aaa],ab,ac],b,c[ca,cb,cc[cca]]]
			],ab,ac],b,c[ca,cb,cc[cca]]]
			RETURN
		,ab,ac],b,c[ca,cb,cc[cca]]] <-------- *2
		ab,ac],b,c[ca,cb,cc[cca]]]
		,ac],b,c[ca,cb,cc[cca]]]
		ac],b,c[ca,cb,cc[cca]]]
		],b,c[ca,cb,cc[cca]]]
		RETURN
	,b,c[ca,cb,cc[cca]]]
	b,c[ca,cb,cc[cca]]]
	,c[ca,cb,cc[cca]]]
	c[ca,cb,cc[cca]]]
	[ca,cb,cc[cca]]]
	RECURSE
		ca,cb,cc[cca]]]
		,cb,cc[cca]]]
		cb,cc[cca]]]
		,cc[cca]]]
		cc[cca]]]
		[cca]]]
		RECURSE
			cca]]]
			]]]
			RETURN
		]]
		RETURN
	]
	RETURN
```

*1 The string the function had when it started that nested recursive calls
*2 The string it should continue on, 
Based on these two notes I have to pass the string from the bottom (child call) to the top, so next to the resulting node that 
comes back from the child I also need to pass the string or the index of where the child has stopped moving over the string,
it is also possible to have the string stored as a global variable that all functions can reference but as a general
rule of thumb, global variables make the code harder to debug and follow. 