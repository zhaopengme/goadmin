import{Z as w,J as y,w as C,K as B,I as x,L as D}from"./element-plus.de1a1f40.js";import{d as F}from"./message.b9ad1923.js";import{_ as L}from"./edit.vue_vue_type_script_setup_true_lang.af8d5433.js";import{d as f,s as T,a1 as R,an as N,o,c as S,Q as d,u,O as i,P as t,X as a,U as m}from"./@vue.a137a740.js";import"./@vueuse.a2c70c0f.js";import"./@element-plus.3660753f.js";import"./lodash-es.7718a1e4.js";import"./dayjs.bd523028.js";import"./axios.d8168cfd.js";import"./async-validator.fb49d0f5.js";import"./@ctrl.fd318bfa.js";import"./@popperjs.36402333.js";import"./escape-html.e5dfadb9.js";import"./normalize-wheel-es.8aeb3683.js";import"./index.fc436213.js";import"./lodash.48927ea5.js";import"./vue-router.9605b890.js";import"./pinia.706f3ea9.js";import"./css-color-function.b4c88e1a.js";import"./color.a9016252.js";import"./clone.73d1916b.js";import"./color-convert.755d189f.js";import"./color-name.e7a4e1d3.js";import"./color-string.e356f5de.js";import"./balanced-match.d2a36341.js";import"./ms.564e106c.js";import"./nprogress.0f0f7ca7.js";import"./vue-clipboard3.4e164ffd.js";import"./clipboard.7c3d630c.js";import"./echarts.098cf7d6.js";import"./zrender.0646a922.js";import"./tslib.60310f1a.js";import"./highlight.js.7165574c.js";import"./@highlightjs.7fc78ec7.js";import"./index.eb574d2d.js";const V=f({name:"shortLetter"}),Et=f({...V,setup($){const p=T(),e=R({loading:!1,lists:[]}),c=async()=>{try{e.loading=!0,e.lists=await F(),e.loading=!1}catch{e.loading=!1}},g=r=>{var s;(s=p.value)==null||s.open(r)};return c(),(r,s)=>{const n=y,_=w,h=C,E=B,v=x,b=N("perms"),k=D;return o(),S("div",null,[d((o(),i(v,{class:"!border-none",shadow:"never"},{default:t(()=>[a(E,{size:"large",data:u(e).lists},{default:t(()=>[a(n,{label:"\u77ED\u4FE1\u6E20\u9053",prop:"name","min-width":"120"}),a(n,{label:"\u72B6\u6001","min-width":"120"},{default:t(({row:l})=>[l.status==1?(o(),i(_,{key:0},{default:t(()=>[m("\u5F00\u542F")]),_:1})):(o(),i(_,{key:1,type:"danger"},{default:t(()=>[m("\u5173\u95ED")]),_:1}))]),_:1}),a(n,{label:"\u64CD\u4F5C","min-width":"120",fixed:"right"},{default:t(({row:l})=>[d((o(),i(h,{type:"primary",link:"",onClick:z=>g(l.alias)},{default:t(()=>[m(" \u8BBE\u7F6E ")]),_:2},1032,["onClick"])),[[b,["setting:storage:edit"]]])]),_:1})]),_:1},8,["data"])]),_:1})),[[k,u(e).loading]]),a(L,{ref_key:"editRef",ref:p,onSuccess:c},null,512)])}}});export{Et as default};
