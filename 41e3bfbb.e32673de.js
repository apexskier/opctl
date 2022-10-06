(window.webpackJsonp=window.webpackJsonp||[]).push([[19],{159:function(e,n,t){"use strict";t.r(n),t.d(n,"frontMatter",(function(){return i})),t.d(n,"metadata",(function(){return l})),t.d(n,"rightToc",(function(){return p})),t.d(n,"default",(function(){return u}));var o=t(1),r=t(9),a=(t(0),t(216)),i={title:"Inputs & Outputs"},l={id:"training/inputs-outputs",title:"Inputs & Outputs",description:"## Inputs",source:"@site/docs/training/inputs-outputs.md",permalink:"/docs/training/inputs-outputs",editUrl:"https://github.com/opctl/opctl/edit/master/website/docs/training/inputs-outputs.md",lastUpdatedBy:"=",lastUpdatedAt:1616514289,sidebar:"docs",previous:{title:"Hello World",permalink:"/docs/training/hello-world"},next:{title:"Run a go service",permalink:"/docs/run-a-go-service"}},p=[{value:"Inputs",id:"inputs",children:[]},{value:"Outputs",id:"outputs",children:[]}],s={rightToc:p},c="wrapper";function u(e){var n=e.components,t=Object(r.a)(e,["components"]);return Object(a.b)(c,Object(o.a)({},s,t,{components:n,mdxType:"MDXLayout"}),Object(a.b)("h2",{id:"inputs"},"Inputs"),Object(a.b)("p",null,"Let's modify the previous simple op to take an input and use it"),Object(a.b)("pre",null,Object(a.b)("code",Object(o.a)({parentName:"pre"},{className:"language-yaml"}),'name: hello-world\n\ndescription: echoes hello followed by a name you provide\n\n# we add the inputs section\ninputs:\n  person: # the name of this input is "person"\n    description: who to greet # the description is "who to greet"\n    string: # the type of this input is string\n      constraints: { minLength: 1 } # it can have a minLength of 1\n\nrun:\n\n  container:\n\n    image: { ref: \'alpine:3.6\' }\n    envVars: { person: $(person) } # we dereference our input "person" and assign its value to an environment variable called "person" inside the container\n    # invoke echo w/ arg "hello $person" - shell will substitute $person with the value of environment variable "person"\n    cmd:\n      - sh\n      - -ce\n      - echo hello $person\n')),Object(a.b)("p",null,"if you run that, you'll be prompted for the input"),Object(a.b)("pre",null,Object(a.b)("code",Object(o.a)({parentName:"pre"},{className:"language-bash"}),'$ opctl run hello-world\n\n-\n  Please provide "person".\n  Description: who to greet\n-\n\n')),Object(a.b)("p",null,'if you type in "you", the container will run and echo out "hello you"'),Object(a.b)("p",null,"Now you may not want to be prompted for the input everytime you run the op. That's why there's several ways to accept input:"),Object(a.b)("ol",null,Object(a.b)("li",{parentName:"ol"},Object(a.b)("inlineCode",{parentName:"li"},"-a")," cli flag: explicitly pass args to op. eg: ",Object(a.b)("inlineCode",{parentName:"li"},"-a NAME1=VALUE1 -a NAME2=VALUE2")),Object(a.b)("li",{parentName:"ol"},Object(a.b)("inlineCode",{parentName:"li"},"--arg-file")," cli flag: reads in a file of args as key=value, in yml format. eg: ",Object(a.b)("inlineCode",{parentName:"li"},'--arg-file="./args.yml"'),". This flag has a default value of ",Object(a.b)("inlineCode",{parentName:"li"},".opspec/args.yml")," i.e. opctl will automatically check for an args file at ",Object(a.b)("inlineCode",{parentName:"li"},".opspec/args.yml")),Object(a.b)("li",{parentName:"ol"},"Environment variables: If you define an environment variable with the same name as an input on the machine you're running opctl on, its value will be supplied as the input's value"),Object(a.b)("li",{parentName:"ol"},Object(a.b)("inlineCode",{parentName:"li"},"default")," property: You can define a ",Object(a.b)("inlineCode",{parentName:"li"},"default")," property for each input, containing a value to assign if no other input method was invoked (cli args or args file)")),Object(a.b)("p",null,"Input sources are checked according to the following precedence:"),Object(a.b)("ul",null,Object(a.b)("li",{parentName:"ul"},"arg provided via -a option"),Object(a.b)("li",{parentName:"ul"},"arg file (since v0.1.19)"),Object(a.b)("li",{parentName:"ul"},"env var"),Object(a.b)("li",{parentName:"ul"},"default"),Object(a.b)("li",{parentName:"ul"},"prompt")),Object(a.b)("h2",{id:"outputs"},"Outputs"),Object(a.b)("p",null,"Let's take that simple op with 1 input and have it provide an output to be used by another op"),Object(a.b)("pre",null,Object(a.b)("code",Object(o.a)({parentName:"pre"},{className:"language-yaml"}),"name: hello-world\n\ndescription: echoes hello followed by a name you provide\n\ninputs:\n  person:\n    description: whom to greet\n    string:\n      constraints: { minLength: 1 }\n\n# we add the outputs section\noutputs:\n  helloperson:\n    description: a string of hello $(person)\n    string: {}\n\nrun:\n  container:\n    files:\n        /output.txt: $(helloperson) # we bind our output to a file that we will create during the container run\n    image: { ref: 'alpine:3.6' }\n    envVars: { person: $(person) } \n    cmd:\n      - sh\n      - -ce\n      - |\n        echo hello $person > /output.txt\n")),Object(a.b)("p",null,"We are now producing an output, let's reference it in another op:\n1. Create a new directory and call it ",Object(a.b)("inlineCode",{parentName:"p"},"caddy"),"\n2. create ",Object(a.b)("inlineCode",{parentName:"p"},"op.yml")," in the ",Object(a.b)("inlineCode",{parentName:"p"},"caddy")," directory, with the below contents"),Object(a.b)("pre",null,Object(a.b)("code",Object(o.a)({parentName:"pre"},{className:"language-yaml"}),"name: caddy\n\ndescription: runs a simple caddy web server that serves a welcome text at http://localthost:8080/\n\ninputs:\n# we need an input of person to pass to the hello-world op when we run it as part of the caddy op\n  person:\n    description: name to greet with welcome text at root of web site\n    string:\n      constraints: { minLength: 1 }\nrun:\n  serial:\n    - op:\n        ref: ../hello-world # here we reference the other op we wrote, hello-world\n        inputs: { person } # we pass our input, person, as input to hello-world\n        outputs: { helloperson } # we add hello-world's output (helloperson) to the scope of this op\n    - container:\n        files:\n            /srv/index.html: $(helloperson) # we dereference helloperson and use its value to populate an index.html file at the default root directory of the caddy image\n        image: { ref: 'abiosoft/caddy' }\n        ports: { '2015' : '8080' } # caddy image listens to 2015 by default, we'd like to serve on 8080 in this example\n")),Object(a.b)("p",null,"and run it"),Object(a.b)("pre",null,Object(a.b)("code",Object(o.a)({parentName:"pre"},{className:"language-bash"}),"$ opctl run -a person=you caddy\n")),Object(a.b)("p",null,'Now if you navigate to http://localhost:8080 in your browser or via curl, you should see the text "hello you"\nAs you make requests to that web server, you should see caddy\'s log in your terminal'),Object(a.b)("p",null,"The above is an example of how ops can reference other ops, and how they can be composed. Note also how we effortlessly and implicitly coerced ",Object(a.b)("inlineCode",{parentName:"p"},"helloperson"),"'s value from a ",Object(a.b)("inlineCode",{parentName:"p"},"string")," into a ",Object(a.b)("inlineCode",{parentName:"p"},"file")," as we mounted ",Object(a.b)("inlineCode",{parentName:"p"},"/srv/index.html")," in the container"))}u.isMDXComponent=!0},216:function(e,n,t){"use strict";t.d(n,"a",(function(){return u})),t.d(n,"b",(function(){return m}));var o=t(0),r=t.n(o);function a(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function i(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);n&&(o=o.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,o)}return t}function l(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?i(Object(t),!0).forEach((function(n){a(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function p(e,n){if(null==e)return{};var t,o,r=function(e,n){if(null==e)return{};var t,o,r={},a=Object.keys(e);for(o=0;o<a.length;o++)t=a[o],n.indexOf(t)>=0||(r[t]=e[t]);return r}(e,n);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(o=0;o<a.length;o++)t=a[o],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(r[t]=e[t])}return r}var s=r.a.createContext({}),c=function(e){var n=r.a.useContext(s),t=n;return e&&(t="function"==typeof e?e(n):l({},n,{},e)),t},u=function(e){var n=c(e.components);return(r.a.createElement(s.Provider,{value:n},e.children))},d="mdxType",b={inlineCode:"code",wrapper:function(e){var n=e.children;return r.a.createElement(r.a.Fragment,{},n)}},h=Object(o.forwardRef)((function(e,n){var t=e.components,o=e.mdxType,a=e.originalType,i=e.parentName,s=p(e,["components","mdxType","originalType","parentName"]),u=c(t),d=o,h=u["".concat(i,".").concat(d)]||u[d]||b[d]||a;return t?r.a.createElement(h,l({ref:n},s,{components:t})):r.a.createElement(h,l({ref:n},s))}));function m(e,n){var t=arguments,o=n&&n.mdxType;if("string"==typeof e||o){var a=t.length,i=new Array(a);i[0]=h;var l={};for(var p in n)hasOwnProperty.call(n,p)&&(l[p]=n[p]);l.originalType=e,l[d]="string"==typeof e?e:o,i[1]=l;for(var s=2;s<a;s++)i[s]=t[s];return r.a.createElement.apply(null,i)}return r.a.createElement.apply(null,t)}h.displayName="MDXCreateElement"}}]);