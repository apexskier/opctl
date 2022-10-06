(window.webpackJsonp=window.webpackJsonp||[]).push([[14],{154:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return c})),r.d(t,"metadata",(function(){return i})),r.d(t,"rightToc",(function(){return l})),r.d(t,"default",(function(){return b}));var n=r(1),a=r(9),o=(r(0),r(216)),c={sidebar_label:"ui",title:"opctl ui"},i={id:"reference/cli/ui",title:"opctl ui",description:"```sh",source:"@site/docs/reference/cli/ui.md",permalink:"/docs/reference/cli/ui",editUrl:"https://github.com/opctl/opctl/edit/master/website/docs/reference/cli/ui.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1604531186,sidebar_label:"ui",sidebar:"docs",previous:{title:"opctl self-update",permalink:"/docs/reference/cli/self-update"},next:{title:"UI",permalink:"/docs/reference/ui"}},l=[{value:"Arguments",id:"arguments",children:[{value:"<code>MOUNT_REF</code> <em>default: <code>.</code></em>",id:"mount_ref-default-",children:[]}]},{value:"Global Options",id:"global-options",children:[]},{value:"Examples",id:"examples",children:[]}],p={rightToc:l},u="wrapper";function b(e){var t=e.components,r=Object(a.a)(e,["components"]);return Object(o.b)(u,Object(n.a)({},p,r,{components:t,mdxType:"MDXLayout"}),Object(o.b)("pre",null,Object(o.b)("code",Object(n.a)({parentName:"pre"},{className:"language-sh"}),"opctl ui [MOUNT_REF=.]\n")),Object(o.b)("p",null,"Open the opctl web UI and mount a reference."),Object(o.b)("h2",{id:"arguments"},"Arguments"),Object(o.b)("h3",{id:"mount_ref-default-"},Object(o.b)("inlineCode",{parentName:"h3"},"MOUNT_REF")," ",Object(o.b)("em",{parentName:"h3"},"default: ",Object(o.b)("inlineCode",{parentName:"em"},"."))),Object(o.b)("p",null,"Reference to mount (either ",Object(o.b)("inlineCode",{parentName:"p"},"relative/path"),", ",Object(o.b)("inlineCode",{parentName:"p"},"/absolute/path"),", ",Object(o.b)("inlineCode",{parentName:"p"},"host/path/repo#tag"),", or ",Object(o.b)("inlineCode",{parentName:"p"},"host/path/repo#tag/path"),")."),Object(o.b)("h2",{id:"global-options"},"Global Options"),Object(o.b)("p",null,"see ",Object(o.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/reference/cli/global-options"}),"global options")),Object(o.b)("h2",{id:"examples"},"Examples"),Object(o.b)("p",null,"Open web UI to current working directory"),Object(o.b)("pre",null,Object(o.b)("code",Object(n.a)({parentName:"pre"},{className:"language-sh"}),"opctl ui\n")),Object(o.b)("p",null,"Open web UI to remote op (github.com/opspec-pkgs/_.op.create#3.3.1)"),Object(o.b)("pre",null,Object(o.b)("code",Object(n.a)({parentName:"pre"},{className:"language-sh"}),"opctl ui github.com/opspec-pkgs/_.op.create#3.3.1\n")))}b.isMDXComponent=!0},216:function(e,t,r){"use strict";r.d(t,"a",(function(){return b})),r.d(t,"b",(function(){return f}));var n=r(0),a=r.n(n);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function c(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?c(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):c(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var p=a.a.createContext({}),u=function(e){var t=a.a.useContext(p),r=t;return e&&(r="function"==typeof e?e(t):i({},t,{},e)),r},b=function(e){var t=u(e.components);return(a.a.createElement(p.Provider,{value:t},e.children))},s="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.a.createElement(a.a.Fragment,{},t)}},m=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,o=e.originalType,c=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),b=u(r),s=n,m=b["".concat(c,".").concat(s)]||b[s]||d[s]||o;return r?a.a.createElement(m,i({ref:t},p,{components:r})):a.a.createElement(m,i({ref:t},p))}));function f(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var o=r.length,c=new Array(o);c[0]=m;var i={};for(var l in t)hasOwnProperty.call(t,l)&&(i[l]=t[l]);i.originalType=e,i[s]="string"==typeof e?e:n,c[1]=i;for(var p=2;p<o;p++)c[p]=r[p];return a.a.createElement.apply(null,c)}return a.a.createElement.apply(null,r)}m.displayName="MDXCreateElement"}}]);