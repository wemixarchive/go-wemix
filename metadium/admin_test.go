// admin_test.go

package metadium

import ()

func TestDistributeRewards() {
	pool := common.HexToAddress("0x1234")
	maint := common.HexToAddress("0xabcd")
	var members []*metaMember

	members = append(members, &metaMember{
		Addr:  common.HexToAddress("0x0001"),
		Stake: big.NewInt(100),
	})
	members = append(members, &metaMember{
		Addr:  common.HexToAddress("0x0002"),
		Stake: big.NewInt(200),
	})
	members = append(members, &metaMember{
		Addr:  common.HexToAddress("0x0003"),
		Stake: big.NewInt(300),
	})

	doit := func(six int, poolAcct, maintAcct *common.Address, ms []*metaMember, amt *big.Int) {
		n := len(ms)
		if poolAcct != nil {
			n++
		}
		if maintAcct != nil {
			n++
		}
		rr := make([]reward, n)
		distributeRewards(six, poolAcct, maintAcct, ms, rr, amt)
		s, _ := json.Marshal(rr)
		fmt.Println(string(s))
	}

	fmt.Println("pool, maint & members set with 101 total")
	doit(0, &pool, &maint, members, big.NewInt(101))
	fmt.Println("pool & maint set with 101 total")
	doit(0, &pool, &maint, nil, big.NewInt(101))
	fmt.Println("pool & members set with 101 total")
	doit(0, &pool, nil, members, big.NewInt(101))
	fmt.Println("pool set with 101 total")
	doit(0, &pool, nil, nil, big.NewInt(101))
	fmt.Println("maint & members set with 101 total")
	doit(0, nil, &maint, members, big.NewInt(101))
	fmt.Println("maint set with 101 total")
	doit(0, nil, &maint, nil, big.NewInt(101))
	fmt.Println("members set with 101 total")
	doit(0, nil, nil, members, big.NewInt(101))
	fmt.Println("none set with 101 total")
	doit(0, nil, nil, nil, big.NewInt(101))

	v := "1000000000000000000000000000000000001"
	ba, _ := new(big.Int).SetString(v, 0)
	fmt.Println("pool, maint & members set with", v)
	doit(0, &pool, &maint, members, ba)
}

// EOF
