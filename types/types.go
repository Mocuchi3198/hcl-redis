package types

type Create struct {
	Name string `hcl:"name,label"`
	Type string `hcl:"type"`
	Value *string `hcl:"value,optional"`
	Fields map[string]string `hcl:"fields,optional"`
	Values []string `hcl:"values,optional"`
}

type Assign struct {
	Name string `hcl:"name,label"`
	Fields map[string]string `hcl:"fields"`
}

type Add struct {
	Name string `hcl:"name,label"`
	Values []string `hcl:"values"`
}

type Get struct {
    Name  string  `hcl:"name"`
    Field *string `hcl:"field,optional"`
}

type Delete struct {
    Name  string  `hcl:"name"`
    Field *string `hcl:"field,optional"`
}

type Document struct {
	Create []Create `hcl:"create,block"`
	Assign []Assign `hcl:"assign,block"`
	Add []Add `hcl:"add,block"`
	Get []Get `hcl:"get,block"`
	Delete []Delete `hcl:"delete,block"`
}