package visor

import "github.com/spaco/spo/src/coin"

const (
	// Maximum supply of spo tokens
	MaxCoinSupply uint64 = 28e8 // 2800,000,000 million

	// Number of distribution addresses
	DistributionAddressesTotal uint64 = 280

	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal

	// Initial number of unlocked addresses
	InitialUnlockedCount uint64 = 25

	// Number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 5

	// Unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 60 * 60 * 24 * 365 // 1 year
)

func init() {
	if MaxCoinSupply%DistributionAddressesTotal != 0 {
		panic("MaxCoinSupply should be perfectly divisible by DistributionAddressesTotal")
	}
}

// Returns a copy of the hardcoded distribution addresses array.
// Each address has 10,000,000 coins. There are 280 addresses.
func GetDistributionAddresses() []string {
	addrs := make([]string, len(distributionAddresses))
	for i := range distributionAddresses {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// Returns distribution addresses that are unlocked, i.e. they have spendable outputs
func GetUnlockedDistributionAddresses() []string {
	// The first InitialUnlockedCount (30) addresses are unlocked by default.
	// Subsequent addresses will be unlocked at a rate of UnlockAddressRate (5) per year,
	// after the InitialUnlockedCount (30) addresses have no remaining balance.
	// The unlock timer will be enabled manually once the
	// InitialUnlockedCount (30) addresses are distributed.

	// NOTE: To have automatic unlocking, transaction verification would have
	// to be handled in visor rather than in coin.Transactions.Visor(), because
	// the coin package is agnostic to the state of the blockchain and cannot reference it.
	// Instead of automatic unlocking, we can hardcode the timestamp at which the first 30%
	// is distributed, then compute the unlocked addresses easily here.

	addrs := make([]string, InitialUnlockedCount)
	for i := range distributionAddresses[:InitialUnlockedCount] {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// Returns distribution addresses that are locked, i.e. they have unspendable outputs
func GetLockedDistributionAddresses() []string {
	// TODO -- once we reach 30% distribution, we can hardcode the
	// initial timestamp for releasing more coins
	addrs := make([]string, DistributionAddressesTotal-InitialUnlockedCount)
	for i := range distributionAddresses[InitialUnlockedCount:] {
		addrs[i] = distributionAddresses[InitialUnlockedCount+uint64(i)]
	}
	return addrs
}

// Returns true if the transaction spends locked outputs
func TransactionIsLocked(inUxs coin.UxArray) bool {
	lockedAddrs := GetLockedDistributionAddresses()
	lockedAddrsMap := make(map[string]struct{})
	for _, a := range lockedAddrs {
		lockedAddrsMap[a] = struct{}{}
	}

	for _, o := range inUxs {
		uxAddr := o.Body.Address.String()
		if _, ok := lockedAddrsMap[uxAddr]; ok {
			return true
		}
	}

	return false
}

var distributionAddresses = [DistributionAddressesTotal]string{
	"iLjCBASLoTPWeZ6sWE3MwTMCQQZDQTVGCp",
	"2axuLZLe4dbZHWnjgrX5snVQ6u9BgBeXFGC",
	"2DPqCjeeFgKAag3XxXdC8ENzeA3wMet4sFV",
	"2N5cfZanuWgtgfNSTtyF1eJyc67zJ1fDYoH",
	"Y8Rb3mqbibFR8tBktYchTyWi2cSbRxp5X1",
	"Hm3buixqBpakRZ3o6dEVUuU8cv2ygyJ7YL",
	"kEDdtR9Cs5wH8WEpWjumw7PPuemArTVioE",
	"D1Stt4yVzCuVs7o7WweuKbyapPrYTud3xQ",
	"2ZVhn5Ga9decYc8FcKj8CPLr8Khe4AzU7C8",
	"2ehpX1Jp1aQYvkytqeTHRXnhJtKZ9i6wHWx",
	"67UWKrK1EpYCYTuwo7CmWMVypCTs9tGLKj",
	"264MBmEri5GfW13A6Pc32CenduqX8vQpCU6",
	"2f2XRHVcce6RoCU78M67qSdRscdtx7TS3hX",
	"6nMu8RdMNY69yP2ECAhpGsq1d2EMBtjqhu",
	"2H7mnjsEPBrQG4J13nXhYXG3e9Mn6iHCpdZ",
	"2huNdKNbE1R33bQea123xoJc6GVAFaJmc9N",
	"tAfRU7qh4oZop99KMVYHKCuVwTrX7ama3H",
	"2Q8t93L5NMcf6qMvWvREyHNryuwaucUS5Um",
	"ZmfWZtmbfSVYys1kNMQNCAakhpu78KSXVU",
	"2Hm3PB5SSnmvDLUJU4oMSJeXNeCXMyuoiFh",
	"PkkVTi995UTiBUa5vL2zAvx45vngNKRVFp",
	"oGCJax1iKD8FCyNPNf1DdsMFUb3WGSnYnX",
	"PzYiPR9roYyp96fgTLQDxUAAX8v8HuxY56",
	"McuKFF7rtxc7bVRn9tAd3Krsqsrj9Mp7Hf",
	"23ECH6u7YSynoY2599xp479ENkPU1yVHZU4",
	"JUPnZCGo2qKZ4E29knG85jggtEfG6SQkkE",
	"2Fq4e7zy5rEB5RbgD9uFfH8G8zt5m34npmK",
	"ay8b32xZSLpLS4bDgTxg9dPj7oMTgWW9aP",
	"2AUsX3TJdfS6zpQjcpeR6y6hra1YJ6XBgiS",
	"VEHMi2Nhv31ftrwScTm2bkky9iWXRQspEj",
	"SZbtNNyXCqGVNXuBf98S2jodWku1mWvTrQ",
	"2h6yUxfNw7ET98DC8ALd87j8mqtJ5gYKSPQ",
	"2Yt5QW5o3M3uyQbPfodA4Kj3wq8WsuF6TAV",
	"Lnt9BHXMxuumubkRQzWnS6s5Z9cToXdhA5",
	"2avR8yznSJ8VEyBUunz9gzbynPm2w1NALVH",
	"2aMYvpvmR5Cg9GaV6sR5mpeX2NqGRGFPoE",
	"x39sP6gXvrkekguN8ZngHFasw355Dbid5F",
	"2hGKQkt9gpudKedpsHujjuwLnZfCWYkb5Pd",
	"2Pms7CAb9EyasxnYuRGQE4dseYaYyMkUTMe",
	"kPPAaXG1fzu7mA5qWBhRVBuLmEzyX7WVdJ",
	"Vt17JaBkGBWi97291fe9GADgurXf28ek4u",
	"wjQppoFEgLy1EBZt8NzcGvqC58MasZsWqo",
	"ZUQuS6MzoC3JZwtreh7C9bMu5BxTwWtHpQ",
	"Hqs3r9nxmR3hSbitw5mcSyZvepwg1i8Nvm",
	"8Nmes4y52URhSfEUWrPKV15mZKkTjDvjke",
	"bsdoSDB7deBhvRMWQB9TgKeDpUXpfpNeq",
	"2Ar1xmpAH9dn9sXS7j7GiY7YhxGRrHpcUib",
	"WJuNEvRTfqVcSV7hHjrdjcKsjXhCNejAEu",
	"T3UyUH9gK1KxPcjv9m93LcPv6FKsJmK617",
	"2VtekN7JepQYkZCeBAQ1wpm8ghi31UDWoha",
	"2D3ws337bn4uaK7TcMjLcPxaBmcEfFZzRvr",
	"2XMSAxo8xpAKvCsyLKjgN7uC1582VD3Hx2y",
	"2LpNd18gu3ABf6HLELEqxFysq3ojmGCACF9",
	"2fQ1VQmG5mAYrbjG2wtLN78RJBdy6ZF7hm",
	"2T2LPTSTB2h6sK1bEpjKT8itX7R3kTkUtGt",
	"2Dpb8NuvYWhTZFHsQw3NXsq854zFExjMEnE",
	"2K1Ga4atyqZWojxR1LA4znZJL4ThDq2WhSM",
	"YWujJW4vVTFPVpagXHnWYsPjPxpPqevxzv",
	"2Y8h9mSEXMCtFwJSYfNbh3LRpFhhqKcZDzC",
	"ghEchATpaX2g4gAX83JJsrdHBvWz6XmSun",
	"2Yei2pFswwvoGrNrNtjmDbceUYbxjzG4g4L",
	"2asX1Z3KncsTMXNwCQkTWhZi4KLUQBGSumW",
	"PkMPwG71nzyEnmrTzryLcynDnxUKwYaDLR",
	"2Vodo5Rrm7Y1PT3Nn4KkCpeaoi27JdyfaRx",
	"uWSdLi6KAWTpvZyhfeNCH85CD3w8HE8MJV",
	"2RVkwWSiN9kBfXFThQvm6t48fe1Nv5qfoFg",
	"2sabep2Vm8NAp4TJi8pzpQaqBbpbiWFo7e",
	"2SRhcEvMXzKxii9NiysRTit56eBJc2Drtn5",
	"KoMbXdRs7f33XRB7mFGfXyDmG2CbrEoZcg",
	"f2Bwjr4m9DEm5SzXz7uJAkHCq9barKCs4a",
	"27hxCVwANGhqNsEBwcDkZ3cCuFUP7hsuKY8",
	"hPXK8US67DFbXcimUHenLNzEhpMMeDoFjL",
	"2Zbw2VtrZxNDG4xBf9Lawy1T6prcqvYKjGx",
	"27m9nahQsAAiL3LXVDA6TFtthYRwm99KfZ8",
	"mdmmXBjfoRQcxp9YwoF9ToVgG2Lg3QgpVx",
	"2kXcfvuaaMMw1tFUPtuaas1GVpuVBJpjHMS",
	"tgKJErYnNBFGMfU9kYJRbJjJPSBh6oZZ7j",
	"212pLpgzdBXymygwnCW5X66DCaZv8fGRoWB",
	"27BvG4Ys8M3mCy8iQvgQY1jVtbBeW4KkuQc",
	"2EYjsCSaLaBe678y9hMbagJCZXi1JtfVqoo",
	"MCX3ehR52o1mbf8HrTDYacS5rQms59uhPx",
	"2XeFKGdDPpAcRL6n35fr3kMdCXL6iyQbiyr",
	"g3Lj6ciDWzsTjJmMx1A6cER5DE6SYEFh3x",
	"QPpagG24UYerP6yiDegdSTfXYtqFaoquha",
	"u2iXtAyF1Fr7t926aA2LaQSdrpUNRT6xbz",
	"2ixSE7WKdrxHxcRywPF9DqSXPvMsigX8UXX",
	"octRejftzSbwpnXHiAiTvLgwBgvFcfJwsx",
	"2TwVdEdXjkJ7gLhtGZfpRUFVwpsE6nZctct",
	"2Aao4awkomFsBR8aZvQSUNiVMQchP2rcExU",
	"2b4nr5dTEoRDcgB6ZpUUyoTGwAwjvptbqrY",
	"23sqMkifJyH9aXGwggCFfhpx2jwoqRFLxY7",
	"MPshSMyRBcSuQ3xdaUHfkLnPstNDY5SXjm",
	"qwuFhacVCgEMdXauRfwZEoxvDmiJowdaoU",
	"MfGdAu9Jbvkr6oJus79jMeScP9tbykkATt",
	"2AWj9NBpgmtd6qQR1VWd2UGV7BaQtSnk8nE",
	"2bv2CxScxrSu719SZ6R7MemAcE9Z3bT8dbK",
	"ArqnvmRXPxS2KkuBNriv8nqjusMTbEmeAe",
	"xcNCxGWUXLGmFUjspe7LKyLSEkxDaFf6d1",
	"2QdYdW6ob1Dh2hTLF9r57LBFemQ8ur32DBg",
	"yYoyGM9noeScwhbE78CEPFz6MTTZK7AFLj",
	"ZyDvyMarfWDm8qBYWMQKFBSLpffViu8AYX",
	"28mFfJMDoBmG7oiAcm4kWfP7CRBki5zyuH2",
	"7aqdA6fPLvvYpVr66ocUjg9Anktk6VLTTh",
	"XVmstZTpwYfSzg3xf56fGGdiqfg6bBZPXz",
	"2PBvETLiQU88GE6h8zkT28hHqpJHhHtH8zp",
	"3wi9eGm9Fk6yS7ZqWjbM6aUmWDW2zFkHGm",
	"2C7FYyeBS9vSB97hZaAEGyg7eLH3UZUwQ3d",
	"s1sYcna9yaevX72aARw9YNyCSTkNEYXkEx",
	"41LA63Kw48qb3TawcPXXs4S3vQMS2HDkgx",
	"GLDL3BrLepEHviiM6DdSWgdjpPHUcquP8E",
	"2kFZVsUN9tLSV1CWHJhCp3jQutXSYnW8HVS",
	"2ZRk2wa7LEmaGquvbQV9212u2KdaZW4gnVk",
	"2FX49mAuLi2a7Qa1eLLVMfEBrwuVRz2xMZJ",
	"SfoQauhnHhVZ2HnJENs9uLPK1SSKoTv7r9",
	"UxBmJSWFSpZad72ijXaEQwxU7fWEUsyVhT",
	"ce7kns4QnFCtCJWSdUjYVq57HjSPYUCtTG",
	"2QQfVPpi6ZAVpnzT9icETzwxkTKHRCFt89B",
	"tkUyhNCG3xKb959vt5AzME8JjyjaNyraet",
	"266NZHZEhizN6T5Q8BUHCtvQtSX38nNxaDT",
	"2bjryyfdbSb5otbnHQW6H4meeCWTHDCfNgp",
	"aYefW3okmcY3onDfrhNCBzhhavpnPFf3q5",
	"cVeUGcnkJRgzDSsndnWrFDQVpW8LQ3Rp7R",
	"2Wg9MEhUfDaXV8pC9gQXzpu8ULnTG7dk9Yi",
	"nBBhWVUwfWvTwGePgr6W7dBxLegQz2srMj",
	"2FiDiF3AZxW9F9JkaPT2sNrt3tKRX2A7ygS",
	"mSWoSHKxNQr6eWKfBXFf69PWCg6xch8V7H",
	"RL31cfomPKcTsgJP4HLntz2Eq3w5FLBrnW",
	"ACuprBKdnkgg6aSEb3ESfx15jgiEcagDwk",
	"2TLashKSR5dqh2K34yaJBCE5MWLLeV2W67U",
	"ezNsiff1ZD8uFDapPPSGwwA3af95pyyBWP",
	"Ra3V4FRv5m82ijYV4DPdozArXoS3pmCEab",
	"KNhWU3qNQ3Jkc8L5vzVuVm7uvGh6SyDPVq",
	"mz1oGsvKNuCydrTXBAAvTGCmyUKWGwUpcY",
	"JnEQRS2QwkNxKcaFQJfYoCLmUMDhYEuh7s",
	"25ebvTaiEaSpnfUMNu5Qf5BMHb6ADET7sa9",
	"2ARMUMuY33wCBvZk6zBmUJgzxUVVk5FyFaf",
	"4u67hpHPTHWGT9dwptEHL7MdgQ8jxC2ZWj",
	"wQH4wbEjDfuz1hAvtmWoB2KjA2tRnmgynf",
	"2Kts6g8dyV8UdoXBCyWg97iV8um8RXxfE8R",
	"Kmfiw4fmNGDoC8cshya8ySDgxCXu8txdyD",
	"wNRFZ7B6mcQ5uZ2f2ywzhPbnSpXW786MN6",
	"3Y7L5ZBf6UdJvdnDgRHrs5LnGVNY5HJFtG",
	"NZBHwbEHXa5AoSJMtRdee6PxTac1mmpKPg",
	"rbdDCrcV53SnJoheNgS3Q7RRao1wRjcHsR",
	"2Wpy4G9SVGnE1BAoZXZcJ8gZWAuVsmSpUcb",
	"7AoTa786NTSi3FC9Pfi3FGQ7Z7ERxUrd5Y",
	"24QqhGj5pFMHr7yEZ8BvEY94qbhkGPddKmS",
	"2YjQgcpLS7KcL6y1FEqx1VGBmqARxXw8NRf",
	"2bqSXZah7tpXwnBDieqgzgm3odKzrivJtgf",
	"2JDZc2MmxB7CZ7tFtEhWsgKYfoD48SwybmN",
	"ejKxdmGVqpeonq2bCMEKLf17tTTVWCpPmd",
	"298LnsYxUqXWUbxqjLMJNNUVJWaNhT1uaEA",
	"AwAB29RHTSEGH2V5bkxAAjmfYUgTuptzhc",
	"2hEEDVTQS5BCre1EGsxHuZEBYiqu2X8jKM9",
	"2Y1J8fWHD5xxpxSGAFa8AmrGCvcw1WuuC6j",
	"2UBxKNXD7PKHVjD5iqoRqikTCAdvZk6PXrt",
	"2Su2nV29eTBuJvWuefQCD3Ypr8oCu14EgQE",
	"2fasNfhsWN4c8EQa18TcRDNCWjsmNDMQk3V",
	"Qx7oqivqREvjPJFn37RjbMjFDiYhdL25Sf",
	"iiGe44h6sYkcmwoQpGf2GFbg2udU3oR7HK",
	"2S8SXYRUnVeHUcynJ93BYyq2B5aaZqy5cAu",
	"2VTxmzERLL1W9p3LKTfcwq1aFg3NEiJ9ka9",
	"Kgif5YbuziFSYEkSHD3ddgivjwJhNcPLtY",
	"ELvHZZ6iB2vNDtDb6CVdSoDHAMmWNRBqzk",
	"2GmLLv4Cu9bNiptHYruq1UD3k6ky6puzpZJ",
	"2ZPjkuoVhB5Y6Hj3SPhvMgSXCfiuHe8CDhg",
	"L2qMrZYLwAAkEXJsWivrw9Xtpdc1ynVTEV",
	"2VgvdMXDnGzWTJ9KXgyouDW3vcpsDmcpZNe",
	"DqgfVwKYJZbiWWeEenBrqKddg6xFwtH92M",
	"2EGBLv2jVMvt8rcxUbzRtkPtV1RnzwRT48u",
	"iDXCSoQvk1qHqfPBmNpKMSvVHCYs91SSVE",
	"EMiEpR2tUvJoVmmYam3Apbxa1tZLw67eKK",
	"4M3WFNkcqR1iLs2gBkbZMHMHpBBrtJmPgW",
	"2CxsUHW989rMHLX8bane4rCxVeUtjZFg2it",
	"2E5uayy2RdwJqkmDN8WkuvigrNiDzMJZvPL",
	"2APaWUpfpbG3yxo88jJXYRxgZRrGNQ5cCsa",
	"2H9oAAmwiuRkZUkQBJV7m97tx1nwKZkjFj5",
	"GE6cqF8Y92LShHyr2wRdDsoDDhHKzJSoBm",
	"B11yLBYYB4zUFfhw3Qu8rWEEWzQxmSegLB",
	"2gcgXt8499VofYtiTCLRkYLVDy6NStmfTft",
	"29vFE4G5rrC7W42gAkAGWxR6NHhxbjjGCSf",
	"2ZGRAC1dkVVTjCpgfajMPiLj7gbmNk4mV7R",
	"bAfKaqYXzwdMnxE3h1Um8DEGG7HQdhLhmN",
	"SJ1yTnq9Zz6iv3ykJcjaLN7cDc9G1v6icG",
	"hU3T5HyQFvUa8bVcqN74xUo41yhpHNgXPq",
	"MHvkJVguwo6a8KFm1kxacPEj8WABqohCAr",
	"gXLBgkcmbPxMJRNfwA5N9wPNK39LAchsDF",
	"25Y3XfAUFG7vbepo3jNYUPn1JCsicCikoAr",
	"wLoVsRcTcpocEYgnTAxa1gPWyLMHhvsCr",
	"cHGNJHDxVScksnK7Fq7sKAa3ENc9wsWEoS",
	"2R8E8EAQ2cDkzutSxWjEezgwaphFA7XF6NV",
	"2Xc8rQqoQNGFhaRt2tYLPtxtrwMzwyM66wh",
	"2ZA8JCMgYxzVQhMcS95bqF2o9Ajnin3zdu3",
	"STYHmRX8FNzE6VwRHqYh2tYm1idtnSQYH1",
	"z3tgu4um8BdC3Fdh9y9yMJSLwAR4tqaryo",
	"25pB49n1ZCpb8f6d8a5kXx4HzVXQij5Q7FR",
	"WwL2kzv7gri7E6MBydb9xu7WxhPkSTRhyG",
	"ZJUXXze66v3j1SVnrAUkkKzMSFmZB1mahr",
	"2ZH6mc1h4y1AXBbg6v2CdLc7Pa2oTUtT8GT",
	"2TUh4ZCVKPBoN6hcx3oxT1aGM29vJhvWb86",
	"8GAf1G3wATSR7SnLjXEyy1tQwunSaiq3V4",
	"Q5nvPwmVnMiobi777Ev49R8MHGPj1PTHbD",
	"94Z1KD7FDogpNFHmHhaYz7TxoSa6MeAy3f",
	"2M2smEBtJkx6nyMSoMQsXDmSvcLDtThvWeh",
	"5RBn79mZp8AaqP9eixjyCLCAqAgg2fLmUk",
	"cfLJEE74iLZ2sopyDYRdEuBbZZZ4Lgr31c",
	"SCeZ4Yvehz62xm1Z3kdx4XiKCBM9ToXqha",
	"Bof6qrpBazie6i856CoVCUSP8NiJzn6UsB",
	"2FMdvA4UfQxy3chj9krJdUBidAr9WgJwFk2",
	"yzz6D15E6WswYQ1WsHfFZqwsVQxYJCr8Rz",
	"VNPH1LQw3PujMpXeS2vgVLyDp7WTUUTmss",
	"2Ash5GUVydRR6J6z9Jo6NCYbQMJHsdY86Wv",
	"gnjtv1tqLSHssE3n5E54PYxQw4Wi8JNGF8",
	"t6NviE8qjXYmvbqSxsvSJtNhZABCUcBGXW",
	"jvuaNzujzLeBCzaJNDhM48vBumcGjQzcLi",
	"22Y6wEWEJT3KGGTJYdM5FgnQZqsDohdx9tc",
	"mB1mfME1QTrBHfU8bAmiWrGHDa6mnPz2ef",
	"9t1YvStLV3XG87kjmGVE7fQFSnMZFM4qd3",
	"2Ax5RZbBExqkoY4zuaEYv8giRix3o355bMA",
	"2YWMvnqAfM7LcJKn8zHBnig4jTnnmwPZ6sw",
	"2L73hdwfvGhreASVN9cG9TAjsfrC7QGdcFf",
	"s9WWDMQjLg1naVEn63LjiRgEKwDibNA3eX",
	"d7k94xoqbX89K1grBttY4kBNsfsrPiTAvJ",
	"2gU5aaqQuZMD7R7f1aqEERVrP1xePqrPjzm",
	"2Ado7Fz6Rwxi5LR7kU2rD74toqEP8H5G4N5",
	"Fv6Pt7G9oCS9KPVmSxT5nSzUwFF5Ki2MTY",
	"JN93ayBnuPpcBw9GVYQsRhBhUfgPjQnFtY",
	"2EpmszFqXFWfy4huWv2J8vhLe48fRgpcfzR",
	"2dkcJwKAJYB1hi9VMoYL6J26YB97voLUJJC",
	"2McsZ18tFWHuq7HdVCzK8c8Ao2zW6tTTsJi",
	"5U98DiqstKMkFuyQPaidSNLApQRVXTSVvz",
	"umYQFVmbAaCDoq3s3WypYr58SuecwFqSNF",
	"2R6ZejyZpGY1Ca1Bm5jxZZjpQTu3XLcmPvG",
	"2UYYmhLoRJu4ywJuSPXTe6tQfmomC3DBzvX",
	"2VHKFt58teiu3uM48pfKffKCJesVDFvJpqt",
	"gPhRENwZJA2Jte4ZRGZrBuoLy1jMvp3wqE",
	"EHHBuEk8J2mC6HhPGa9jza1phPpA1xy3v3",
	"2TJMNZgc1D2GsN9iTk21s52xWDeo1bCzE8y",
	"vVAp9YdPFXf6EFP1UivDkvXzkC2zxiTbUx",
	"cKamNAUCKdY185156EyTMKoHHtbSmiJgup",
	"hr7RLTcQGHf5ZnmbDHU9WEALxdjz5tCuB4",
	"B419sHegutjKjmT7gLTd5o2Gy5mj7zVjWW",
	"zKfcErA5DyYafHFWQtBNsQhaLQG36SkWXK",
	"jSr3vTY6dxzWFfSeazF6CNv8smLv2YrBsP",
	"qfdD4XQHnGmxi4Y13edUkydNwvrqSseVo1",
	"dQ7VGcorsVcRLfG2nxu15Xpim8Xfv1WP2k",
	"5dbf272N2D3HYJEnVhxr5Z7B5JsoDUredc",
	"Xy2axp6Bf8UKDwXrr6WGozCiVrJuwCNTNV",
	"eJff4norJotmUHruftqrDNCYrGbwfXdCp4",
	"2RhFY2VsYZWWPCdARmtsoM683BENro2Xa9g",
	"huKCtgNBxeGpmoFfrKtd6Nv5UBciQ7eFqn",
	"2UmkoDwE8iMpW2tFUR1eBpXh1LuvkEpdZnD",
	"21ubakacKzDY4vbc2H8KGcf5j4etZfCBGch",
	"ooE7duUVFCU11YB3sg7BDtR7kquAaZViF6",
	"riwBKhaff1hyt86DPtoXfBPwxL6RfekCM4",
	"4DoEhvaHUSRA5B9tpERaGgnaamM6q1xGLS",
	"2dKDSrwWQPrgZGQR6qDm2MqMgAcdmck8Kjt",
	"2fVGGzhMb45YR9piXiJ4tYghYyn7E9Neapk",
	"2Pp6FKdEX6SXUdr7o5NRE1W5DhbbXoLzB4c",
	"fdRDkoeVbRJLmMFKTGHopDV3ffcomKQyN2",
	"5QzS4EhNUz5tqrzAvcXxqSr64FGJnWPco7",
	"2Wg1J1S7GBsjVdFxEbFSFtLkMhHYbMjTSR1",
	"ekUbXqnFayhEiDewXTsqYSZq8i8exH4wt3",
	"2g89VdtVaitDGUTPco7GAHx3QQrTiYATCho",
	"2gtFcrxrw1qYEeJBARAPM7epdYay9vzHUhF",
	"kqHsMghHeveveJ6wHhxSuFBB7qMMcQEbyX",
	"8raVmpS9L4pmy83qger3PkenpqaodMziU3",
	"29D1RMFpecY51inbPoohpZFBrGtpfXt1tTp",
	"2dso4tB8YhwrswdnR5QT11mPGaa5k8qwybS",
	"YNrTxDWEt9bBrHs2rQMLkZDWEXEZR3w5Nr",
	"W1DXX26vDoL49M6eznmoX9c1rWqBtNJ9CY",
	"vT7P8RbEvLSohK5EW6gCiiK3qyqZmGNJ9W",
	"9KpU8BWeX2DwDyRxLqJntkTjxrY7njNDg1",
	"6ponFo3HVqVUZup5nWFugDTd9HotnhLRNZ",
	"26s6n1z2kxoLWpcQj4ARZGGKDyYZjVKnYFh",
	"zcuH4ndU5oCa1iMSC4dVGifVyYo53DBq5i",
	"27JJGZAiEy9qEXCWvT48zKafpVgKf7ib3tP",
	"EExf83UsPpfBXxfJV9AAGkmbwZjoJv7axw",
	"2SvGgrWFzrHAAhukzrZNBJmfxb5UESDdZe4",
	"2jJwL8yHurM4WVUNDJrrbKHd4preCSKbx21",
}
