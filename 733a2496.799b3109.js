(window.webpackJsonp=window.webpackJsonp||[]).push([[34],{174:function(e,r,t){"use strict";t.r(r),t.d(r,"frontMatter",(function(){return i})),t.d(r,"metadata",(function(){return p})),t.d(r,"rightToc",(function(){return o})),t.d(r,"default",(function(){return s}));var n=t(1),a=t(9),c=(t(0),t(216)),i={title:"Variable reference"},p={id:"reference/opspec/op.yml/variable-reference",title:"Variable reference",description:"Variable references in an opspec file are how you reference data within your op.",source:"@site/docs/reference/opspec/op.yml/variable-reference.md",permalink:"/docs/reference/opspec/op.yml/variable-reference",editUrl:"https://github.com/opctl/opctl/edit/master/website/docs/reference/opspec/op.yml/variable-reference.md",lastUpdatedBy:"Cameron Little",lastUpdatedAt:1617200902,sidebar:"docs",previous:{title:"Rangeable value",permalink:"/docs/reference/opspec/op.yml/rangeable-value"},next:{title:"Array",permalink:"/docs/reference/opspec/types/array"}},o=[{value:"Escaping",id:"escaping",children:[]},{value:"Filesystem paths",id:"filesystem-paths",children:[]},{value:"Array access",id:"array-access",children:[]},{value:"Object properties",id:"object-properties",children:[]},{value:"Directory paths",id:"directory-paths",children:[]}],l={rightToc:o},b="wrapper";function s(e){var r=e.components,t=Object(a.a)(e,["components"]);return Object(c.b)(b,Object(n.a)({},l,t,{components:r,mdxType:"MDXLayout"}),Object(c.b)("p",null,"Variable references in an opspec file are how you reference data within your op."),Object(c.b)("p",null,"A variable reference takes the form ",Object(c.b)("inlineCode",{parentName:"p"},"$(reference)"),", where ",Object(c.b)("inlineCode",{parentName:"p"},"reference")," can be the ",Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"identifier"}),"name")," of an ",Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/reference/opspec/index#scoping"}),"in scope")," variable, a newly defined variable ",Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"identifier"}),"name"),", or an absolute or relative filesystem path."),Object(c.b)("p",null,"If referencing a variable, ",Object(c.b)("inlineCode",{parentName:"p"},"reference")," can be extended with an array access, object property, or filesystem path, if the type of the variable matches or can be coerced."),Object(c.b)("h2",{id:"escaping"},"Escaping"),Object(c.b)("p",null,"Variable references can be escaped with ",Object(c.b)("inlineCode",{parentName:"p"},"\\"),". Because opspecs are yaml, a second ",Object(c.b)("inlineCode",{parentName:"p"},"\\")," is required to escape the escape when used in a normal string."),Object(c.b)("pre",null,Object(c.b)("code",Object(n.a)({parentName:"pre"},{className:"language-yaml"}),'foo: "\\\\$(wouldBeVariableReference)" # not treated as a variable reference\n')),Object(c.b)("pre",null,Object(c.b)("code",Object(n.a)({parentName:"pre"},{className:"language-yaml"}),"foo: |\n  \\$(wouldBeVariableReference) # not treated as a variable reference\n")),Object(c.b)("h2",{id:"filesystem-paths"},"Filesystem paths"),Object(c.b)("p",null,"Ops have a few filesystem paths in scope automatically."),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},Object(c.b)("inlineCode",{parentName:"li"},"./")," is the the current op's directory, where the ",Object(c.b)("inlineCode",{parentName:"li"},"op.yml")," is located. (the current ",Object(c.b)("inlineCode",{parentName:"li"},"op.yml")," can be accessed with ",Object(c.b)("inlineCode",{parentName:"li"},"$(./op.yml)"),")"),Object(c.b)("li",{parentName:"ul"},Object(c.b)("inlineCode",{parentName:"li"},"../")," is the parent of the current op's directory, (the current ",Object(c.b)("inlineCode",{parentName:"li"},"op.yml")," can be also be accessed with ",Object(c.b)("inlineCode",{parentName:"li"},"$(../op.yml)"),")")),Object(c.b)("h2",{id:"array-access"},"Array access"),Object(c.b)("p",null,Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"../types/array"}),"Array")," items can be referenced with the syntax ",Object(c.b)("inlineCode",{parentName:"p"},"$(reference[index])"),", where ",Object(c.b)("inlineCode",{parentName:"p"},"index")," is the zero based index of the item.\nIf ",Object(c.b)("inlineCode",{parentName:"p"},"index")," is negative, indexing will take place from the end of the array."),Object(c.b)("pre",null,Object(c.b)("code",Object(n.a)({parentName:"pre"},{className:"language-yaml"}),'inputs:\n  someArray:\n    array:\n      default: ["one", 2, [3]]\nrun:\n  op:\n    ref: ../op\n    inputs:\n      input1: $(someArray[0]) # "one"\n      input2: $(someArray[-1][0]) # 3\n')),Object(c.b)("h2",{id:"object-properties"},"Object properties"),Object(c.b)("p",null,Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"../types/object"}),"Object")," properties can be referenced with the syntax ",Object(c.b)("inlineCode",{parentName:"p"},"$(reference.property)")," or ",Object(c.b)("inlineCode",{parentName:"p"},"$(reference[property])"),", where ",Object(c.b)("inlineCode",{parentName:"p"},"property")," is the name of the property."),Object(c.b)("pre",null,Object(c.b)("code",Object(n.a)({parentName:"pre"},{className:"language-yaml"}),"inputs:\n  someObject:\n    object:\n      default:\n        myKey: aValue\n        secondKey:\n          subKey: 2\nrun:\n  op:\n    ref: ../op\n    inputs:\n      input1: $(someObject.myKey) # aValue\n      input2: $(someObject[secondKey][subKey]) # 2\n")),Object(c.b)("h2",{id:"directory-paths"},"Directory paths"),Object(c.b)("p",null,Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"../types/dir"}),"Directory")," entries (child files and directories) can be referenced with the syntax ",Object(c.b)("inlineCode",{parentName:"p"},"$(reference/entry)"),", where ",Object(c.b)("inlineCode",{parentName:"p"},"entry")," is the subdirectory, file, or combination of subdirectories and files."),Object(c.b)("pre",null,Object(c.b)("code",Object(n.a)({parentName:"pre"},{className:"language-yaml"}),"inputs:\n  someDir:\n    dir:\n      default: ./myDir\nrun:\n  op:\n    ref: ../op\n    inputs:\n      input1: $(someDir/file.txt) # the file at ./myDir/file.txt\n      input2: $(someObject/a/b/c) # the file or directory at ./myDir/a/b/c\n")),Object(c.b)("p",null,"The default ",Object(c.b)("inlineCode",{parentName:"p"},"./")," and ",Object(c.b)("inlineCode",{parentName:"p"},"../")," references support these extended path references."),Object(c.b)("pre",null,Object(c.b)("code",Object(n.a)({parentName:"pre"},{className:"language-yaml"}),"run:\n  op:\n    ref: ../op\n    inputs:\n      opfile: $(./op.yml)\n")))}s.isMDXComponent=!0},216:function(e,r,t){"use strict";t.d(r,"a",(function(){return s})),t.d(r,"b",(function(){return m}));var n=t(0),a=t.n(n);function c(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function i(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function p(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?i(Object(t),!0).forEach((function(r){c(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function o(e,r){if(null==e)return{};var t,n,a=function(e,r){if(null==e)return{};var t,n,a={},c=Object.keys(e);for(n=0;n<c.length;n++)t=c[n],r.indexOf(t)>=0||(a[t]=e[t]);return a}(e,r);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(n=0;n<c.length;n++)t=c[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}var l=a.a.createContext({}),b=function(e){var r=a.a.useContext(l),t=r;return e&&(t="function"==typeof e?e(r):p({},r,{},e)),t},s=function(e){var r=b(e.components);return(a.a.createElement(l.Provider,{value:r},e.children))},d="mdxType",u={inlineCode:"code",wrapper:function(e){var r=e.children;return a.a.createElement(a.a.Fragment,{},r)}},f=Object(n.forwardRef)((function(e,r){var t=e.components,n=e.mdxType,c=e.originalType,i=e.parentName,l=o(e,["components","mdxType","originalType","parentName"]),s=b(t),d=n,f=s["".concat(i,".").concat(d)]||s[d]||u[d]||c;return t?a.a.createElement(f,p({ref:r},l,{components:t})):a.a.createElement(f,p({ref:r},l))}));function m(e,r){var t=arguments,n=r&&r.mdxType;if("string"==typeof e||n){var c=t.length,i=new Array(c);i[0]=f;var p={};for(var o in r)hasOwnProperty.call(r,o)&&(p[o]=r[o]);p.originalType=e,p[d]="string"==typeof e?e:n,i[1]=p;for(var l=2;l<c;l++)i[l]=t[l];return a.a.createElement.apply(null,i)}return a.a.createElement.apply(null,t)}f.displayName="MDXCreateElement"}}]);