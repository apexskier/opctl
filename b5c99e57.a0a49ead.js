(window.webpackJsonp=window.webpackJsonp||[]).push([[50],{189:function(e,n,a){"use strict";a.r(n),a.d(n,"frontMatter",(function(){return c})),a.d(n,"metadata",(function(){return i})),a.d(n,"rightToc",(function(){return o})),a.d(n,"default",(function(){return d}));var t=a(1),r=a(9),l=(a(0),a(216)),c={sidebar_label:"Overview",title:"Call"},i={id:"reference/opspec/op.yml/call/index",title:"Call",description:"A call is an object that defines a single call within an ops call graph. Opctl supports several types of calls and several advanced attributes to add logic into your op.",source:"@site/docs/reference/opspec/op.yml/call/index.md",permalink:"/docs/reference/opspec/op.yml/call/index",editUrl:"https://github.com/opctl/opctl/edit/master/website/docs/reference/opspec/op.yml/call/index.md",lastUpdatedBy:"Cameron Little",lastUpdatedAt:1617200902,sidebar_label:"Overview",sidebar:"docs",previous:{title:"op.yml",permalink:"/docs/reference/opspec/op.yml/index"},next:{title:"Container call",permalink:"/docs/reference/opspec/op.yml/call/container/index"}},o=[{value:"Basic propreties",id:"basic-propreties",children:[{value:"<code>description</code>",id:"description",children:[]},{value:"<code>name</code>",id:"name",children:[]}]},{value:"Types of calls properties",id:"types-of-calls-properties",children:[{value:"<code>container</code>",id:"container",children:[]},{value:"<code>op</code>",id:"op",children:[]},{value:"<code>parallel</code>",id:"parallel",children:[]},{value:"<code>parallelLoop</code>",id:"parallelloop",children:[]},{value:"<code>serial</code>",id:"serial",children:[]},{value:"<code>serialLoop</code>",id:"serialloop",children:[]}]},{value:"Logical properties",id:"logical-properties",children:[{value:"<code>if</code>",id:"if",children:[]},{value:"<code>needs</code>",id:"needs",children:[]}]}],p={rightToc:o},s="wrapper";function d(e){var n=e.components,a=Object(r.a)(e,["components"]);return Object(l.b)(s,Object(t.a)({},p,a,{components:n,mdxType:"MDXLayout"}),Object(l.b)("p",null,"A call is an object that defines a single call within an ops call graph. Opctl supports several types of calls and several advanced attributes to add logic into your op."),Object(l.b)("p",null,"The leaves of a call graph are ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"/docs/reference/opspec/op.yml/call/container/index"}),"containers")," that run programs to do work. Each node in the call graph will end once all child containers exit successfully or a single container exits with a failure, which will cause still running containers to be killed."),Object(l.b)("p",null,"Containers and call graph nodes communicate with each other by producing and emitting ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"/docs/reference/opspec/op.yml/index"}),"data")," or by reading and writing shared ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"../types/file.md"}),"files")," and ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"../types/dir.md"}),"directories")," (which are passed by reference)."),Object(l.b)("h2",{id:"basic-propreties"},"Basic propreties"),Object(l.b)("h3",{id:"description"},Object(l.b)("inlineCode",{parentName:"h3"},"description")),Object(l.b)("p",null,"A human friendly description of the parameter, written as a ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"/docs/reference/opspec/op.yml/markdown"}),"markdown string"),"."),Object(l.b)("h3",{id:"name"},Object(l.b)("inlineCode",{parentName:"h3"},"name")),Object(l.b)("p",null,"The name is an ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"/docs/reference/opspec/op.yml/identifier"}),"identifier")," used to identify the call in a UI and used for ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"#needs"}),Object(l.b)("inlineCode",{parentName:"a"},"needs"))," logic in sibling calls."),Object(l.b)("h2",{id:"types-of-calls-properties"},"Types of calls properties"),Object(l.b)("p",null,"Every object in your call graph must declare one of the following properties to define how and what it runs. These can directly run an op or container or can sequence further calls in serial, parallel, or by looping."),Object(l.b)("h3",{id:"container"},Object(l.b)("inlineCode",{parentName:"h3"},"container")),Object(l.b)("p",null,"A ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"/docs/reference/opspec/op.yml/call/container/index"}),"container call")," runs a container."),Object(l.b)("h3",{id:"op"},Object(l.b)("inlineCode",{parentName:"h3"},"op")),Object(l.b)("p",null,"An ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"/docs/reference/opspec/op.yml/call/container/index"}),"op call")," runs an external op."),Object(l.b)("h3",{id:"parallel"},Object(l.b)("inlineCode",{parentName:"h3"},"parallel")),Object(l.b)("p",null,"A parallel call is an array of ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"/docs/reference/opspec/op.yml/call/index"}),"calls")," that are executed concurrently in parallel, with no defined order."),Object(l.b)("h3",{id:"parallelloop"},Object(l.b)("inlineCode",{parentName:"h3"},"parallelLoop")),Object(l.b)("p",null,"A ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"parallel-loop"}),"parallel loop call")," defines a call executed multiple times in parallel."),Object(l.b)("h3",{id:"serial"},Object(l.b)("inlineCode",{parentName:"h3"},"serial")),Object(l.b)("p",null,"A serial call is an array of ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"/docs/reference/opspec/op.yml/call/index"}),"calls")," that are executed serially, one after another."),Object(l.b)("h3",{id:"serialloop"},Object(l.b)("inlineCode",{parentName:"h3"},"serialLoop")),Object(l.b)("p",null,"A ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"serial-loop"}),"serial loop call")," defines a call executed multiple times repeatedly."),Object(l.b)("h2",{id:"logical-properties"},"Logical properties"),Object(l.b)("p",null,"Calls can optionally define additional properties to introduce conditional logic into their execution."),Object(l.b)("h3",{id:"if"},Object(l.b)("inlineCode",{parentName:"h3"},"if")),Object(l.b)("p",null,"An if property on a call is an array of ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"/docs/reference/opspec/op.yml/predicate"}),"predicates"),". If all predicates are true, the call is executed, otherwise it is skipped."),Object(l.b)("pre",null,Object(l.b)("code",Object(t.a)({parentName:"pre"},{className:"language-yaml"}),"if:\n  - eq: [true, $(test)]\n  - exists: $(value)\n")),Object(l.b)("h3",{id:"needs"},Object(l.b)("inlineCode",{parentName:"h3"},"needs")),Object(l.b)("p",null,Object(l.b)("inlineCode",{parentName:"p"},"needs")," allows introducing dependencies between ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"#parallel"}),"parallel")," calls. The value of this property is an array of sibling ",Object(l.b)("a",Object(t.a)({parentName:"p"},{href:"#name"}),"names"),". When all calls that need a given name end, the named call will be killed."),Object(l.b)("p",null,"Needs cannot be used in parents, children, or cousin calls (they must be within the same ",Object(l.b)("inlineCode",{parentName:"p"},"parallel")," call). Opctl will ignore needs that don't satisfy this limitation."),Object(l.b)("pre",null,Object(l.b)("code",Object(t.a)({parentName:"pre"},{className:"language-yaml"}),"description: |\n  `systemUnderTest` will be shutdown after 1 second because it's no longer needed by the second container.\nrun:\n  parallel:\n    - name: systemUnderTest\n      container:\n        image: {ref: alpine}\n        cmd: [sleep, 100000]\n    - container:\n        image: {ref: alpine}\n        cmd: [sleep, 1]\n      needs:\n        - systemUnderTest\n")))}d.isMDXComponent=!0},216:function(e,n,a){"use strict";a.d(n,"a",(function(){return d})),a.d(n,"b",(function(){return m}));var t=a(0),r=a.n(t);function l(e,n,a){return n in e?Object.defineProperty(e,n,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[n]=a,e}function c(e,n){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var t=Object.getOwnPropertySymbols(e);n&&(t=t.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),a.push.apply(a,t)}return a}function i(e){for(var n=1;n<arguments.length;n++){var a=null!=arguments[n]?arguments[n]:{};n%2?c(Object(a),!0).forEach((function(n){l(e,n,a[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):c(Object(a)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(a,n))}))}return e}function o(e,n){if(null==e)return{};var a,t,r=function(e,n){if(null==e)return{};var a,t,r={},l=Object.keys(e);for(t=0;t<l.length;t++)a=l[t],n.indexOf(a)>=0||(r[a]=e[a]);return r}(e,n);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(t=0;t<l.length;t++)a=l[t],n.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(r[a]=e[a])}return r}var p=r.a.createContext({}),s=function(e){var n=r.a.useContext(p),a=n;return e&&(a="function"==typeof e?e(n):i({},n,{},e)),a},d=function(e){var n=s(e.components);return(r.a.createElement(p.Provider,{value:n},e.children))},b="mdxType",u={inlineCode:"code",wrapper:function(e){var n=e.children;return r.a.createElement(r.a.Fragment,{},n)}},f=Object(t.forwardRef)((function(e,n){var a=e.components,t=e.mdxType,l=e.originalType,c=e.parentName,p=o(e,["components","mdxType","originalType","parentName"]),d=s(a),b=t,f=d["".concat(c,".").concat(b)]||d[b]||u[b]||l;return a?r.a.createElement(f,i({ref:n},p,{components:a})):r.a.createElement(f,i({ref:n},p))}));function m(e,n){var a=arguments,t=n&&n.mdxType;if("string"==typeof e||t){var l=a.length,c=new Array(l);c[0]=f;var i={};for(var o in n)hasOwnProperty.call(n,o)&&(i[o]=n[o]);i.originalType=e,i[b]="string"==typeof e?e:t,c[1]=i;for(var p=2;p<l;p++)c[p]=a[p];return r.a.createElement.apply(null,c)}return r.a.createElement.apply(null,a)}f.displayName="MDXCreateElement"}}]);