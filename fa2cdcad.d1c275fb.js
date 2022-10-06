(window.webpackJsonp=window.webpackJsonp||[]).push([[71],{211:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return c})),r.d(t,"metadata",(function(){return i})),r.d(t,"rightToc",(function(){return p})),r.d(t,"default",(function(){return s}));var n=r(1),a=r(9),o=(r(0),r(216)),c={title:"Number parameter"},i={id:"reference/opspec/op.yml/parameter/number",title:"Number parameter",description:"An object defining a [number](../../types/number.md) parameter.",source:"@site/docs/reference/opspec/op.yml/parameter/number.md",permalink:"/docs/reference/opspec/op.yml/parameter/number",editUrl:"https://github.com/opctl/opctl/edit/master/website/docs/reference/opspec/op.yml/parameter/number.md",lastUpdatedBy:"Cameron Little",lastUpdatedAt:1617200902,sidebar:"docs",previous:{title:"File parameter",permalink:"/docs/reference/opspec/op.yml/parameter/file"},next:{title:"Object parameter",permalink:"/docs/reference/opspec/op.yml/parameter/object"}},p=[{value:"Properties",id:"properties",children:[{value:"<code>constraints</code>",id:"constraints",children:[]},{value:"<code>default</code>",id:"default",children:[]},{value:"<code>isSecret</code>",id:"issecret",children:[]}]}],l={rightToc:p},u="wrapper";function s(e){var t=e.components,r=Object(a.a)(e,["components"]);return Object(o.b)(u,Object(n.a)({},l,r,{components:t,mdxType:"MDXLayout"}),Object(o.b)("p",null,"An object defining a ",Object(o.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/reference/opspec/types/number"}),"number")," parameter."),Object(o.b)("h2",{id:"properties"},"Properties"),Object(o.b)("h3",{id:"constraints"},Object(o.b)("inlineCode",{parentName:"h3"},"constraints")),Object(o.b)("p",null,"A ",Object(o.b)("a",Object(n.a)({parentName:"p"},{href:"https://tools.ietf.org/html/draft-wright-json-schema-00"}),"JSON Schema v4 object")," defining constraints on the value."),Object(o.b)("h3",{id:"default"},Object(o.b)("inlineCode",{parentName:"h3"},"default")),Object(o.b)("p",null,"A literal number used as the default value of the variable created by the parameter."),Object(o.b)("h3",{id:"issecret"},Object(o.b)("inlineCode",{parentName:"h3"},"isSecret")),Object(o.b)("p",null,"A boolean indicating if the value of the parameter is secret. This will cause it to be hidden in UI's. "))}s.isMDXComponent=!0},216:function(e,t,r){"use strict";r.d(t,"a",(function(){return s})),r.d(t,"b",(function(){return f}));var n=r(0),a=r.n(n);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function c(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?c(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):c(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function p(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var l=a.a.createContext({}),u=function(e){var t=a.a.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):i({},t,{},e)),r},s=function(e){var t=u(e.components);return(a.a.createElement(l.Provider,{value:t},e.children))},b="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.a.createElement(a.a.Fragment,{},t)}},m=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,o=e.originalType,c=e.parentName,l=p(e,["components","mdxType","originalType","parentName"]),s=u(r),b=n,m=s["".concat(c,".").concat(b)]||s[b]||d[b]||o;return r?a.a.createElement(m,i({ref:t},l,{components:r})):a.a.createElement(m,i({ref:t},l))}));function f(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var o=r.length,c=new Array(o);c[0]=m;var i={};for(var p in t)hasOwnProperty.call(t,p)&&(i[p]=t[p]);i.originalType=e,i[b]="string"==typeof e?e:n,c[1]=i;for(var l=2;l<o;l++)c[l]=r[l];return a.a.createElement.apply(null,c)}return a.a.createElement.apply(null,r)}m.displayName="MDXCreateElement"}}]);