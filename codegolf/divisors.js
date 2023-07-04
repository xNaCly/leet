write = (...e) => process.stdout.write(e.join(" ")+" ");
print = (e) => console.log(e || "");

for(i=1;i<101;i++){for(j=1;j<101;j++)i%j==0&&write(j);print()}
