package domain

type Statistics struct {
	Statistic Statistic
}

type Statistic []struct {
	Kind          string       		 	`json:"kind"`
	Type          string           		`json:"type"`
	Unit      	  string              	`json:"unit"`
	Values        *Values        		`json:"values"`
	Hosts         *Hosts   		   		`json:"host"`
	Name      	  string              	`json:"name"`
	Description   string              	`json:"description"`
	Href      	  string              	`json:"href"`
	Id      	  string              	`json:"id"`
}

type Values struct {
	Value   []Value
}

type Value struct{
	Datum  float64    	`json:"datum"`
}

type Hosts struct {
	Href   string 		`json:"href"`
	Id     string 		`json:"id"`
}

