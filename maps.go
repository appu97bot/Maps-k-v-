package main;
import "log"
type animals struct{
name,colour string;
weight float32;
}
func main(){
a:=make(map[string]animals);
a["i"]=animals{"Tiger","Red",120};
a["j"]=animals{"Lion","Brown",150.5};
a["k"]=animals{"Deer","Brown",60.9};
a["l"]=animals{"Ox","white",89.9};
a["m"]=animals{"Dog","Black",48.2};
a["n"]=animals{"jirafi","Yellow",97.6};
for b,c:=range a{
log.Println("Animals list are:",b,c);
}
delete(a,"j");
log.Println("Deleted elements:",a);
a["i"]=animals{name:"Tiggy"};
a["m"]=animals{colour:"white"};
log.Println("Changed elements are:",a["i"],a["m"]);
}

