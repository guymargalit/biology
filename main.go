package main

import (
	. "biology/cell"

	"fmt"
	// "context"
    // "github.com/go-redis/redis/v8"
)

//var ctx = context.Background()

func main() {
    /*
	rdb := redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := rdb.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }*/

    // Approximate ICF Composition Values (mEq/L)
    var icf = map[string]Ion{}
    icf["Na"] = Ion{Charge: 1, Concentration: 14.0}
    icf["K"] = Ion{Charge: 1, Concentration: 120.0}
    icf["Ca"] = Ion{Charge: 2, Concentration: 1e-4}
    icf["Cl"] = Ion{Charge: -1, Concentration: 10.0}
    icf["HCO3"] = Ion{Charge: -1, Concentration: 10.0}
    icf["H"] = Ion{Charge: +1, Concentration: 40e-6}

    // Approximate ECF Composition Values (mEq/L)
    var ecf = map[string]Ion{}
    ecf["Na"] = Ion{Charge: 1, Concentration: 140.0}
    ecf["K"] = Ion{Charge: 1, Concentration: 4.0}
    ecf["Ca"] = Ion{Charge: 2, Concentration: 2.5}
    ecf["Cl"] = Ion{Charge: -1, Concentration: 105.0}
    ecf["HCO3"] = Ion{Charge: -1, Concentration: 24.0}
    ecf["H"] = Ion{Charge: +1, Concentration: 80e-6}

    var ions = []string{"Na","K","Ca","Cl","HCO3","H"}

    v := 90e-15 // Red blood cell (L)
	c := Cell{
        Membrane: Membrane{
            Area: 1,
            Thickness: 1,
            Radius: 1,
            Viscosity: 1,
            ICF: CF{
                Ions: icf,
                Volume: v,
            },
            ECF: CF{
                Ions: ecf,
                Volume: v,
            },
            Ions: ions,
        },
    }

	c.Init()

	fmt.Println(c)
}
