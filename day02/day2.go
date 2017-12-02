package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const rowDifference = 1
const rowDivisible = 2

func main() {
	var test = `5	1	9	5
	7	5	3
	2	4	6	8`
	var result = calculateChecksum(test, rowDifference)
	fmt.Printf("test result = %v\n", result)

	var input = `5806	6444	1281	38	267	1835	223	4912	5995	230	4395	2986	6048	4719	216	1201
	74	127	226	84	174	280	94	159	198	305	124	106	205	99	177	294
	1332	52	54	655	56	170	843	707	1273	1163	89	23	43	1300	1383	1229
	5653	236	1944	3807	5356	246	222	1999	4872	206	5265	5397	5220	5538	286	917
	3512	3132	2826	3664	2814	549	3408	3384	142	120	160	114	1395	2074	1816	2357
	100	2000	112	103	2122	113	92	522	1650	929	1281	2286	2259	1068	1089	651
	646	490	297	60	424	234	48	491	245	523	229	189	174	627	441	598
	2321	555	2413	2378	157	27	194	2512	117	140	2287	277	2635	1374	1496	1698
	101	1177	104	89	542	2033	1724	1197	474	1041	1803	770	87	1869	1183	553
	1393	92	105	1395	1000	85	391	1360	1529	1367	1063	688	642	102	999	638
	4627	223	188	5529	2406	4980	2384	2024	4610	279	249	2331	4660	4350	3264	242
	769	779	502	75	1105	53	55	931	1056	1195	65	292	1234	1164	678	1032
	2554	75	4406	484	2285	226	5666	245	4972	3739	5185	1543	230	236	3621	5387
	826	4028	4274	163	5303	4610	145	5779	157	4994	5053	186	5060	3082	2186	4882
	588	345	67	286	743	54	802	776	29	44	107	63	303	372	41	810
	128	2088	3422	111	3312	740	3024	1946	920	131	112	477	3386	2392	1108	2741`
	var checksum = calculateChecksum(input, rowDifference)
	fmt.Printf("checksum = %v\n", checksum)
	fmt.Println()

	var test2 = `5	9	2	8
	9	4	7	3
	3	8	6	5`
	var result2 = calculateChecksum(test2, rowDivisible)
	fmt.Printf("test result = %v\n", result2)

	var checksum2 = calculateChecksum(input, rowDivisible)
	fmt.Printf("checksum2 = %v\n", checksum2)
}

func calculateChecksum(input string, algorithm int) int {
	var sum = 0

	for _, str := range strings.Split(input, "\n") {
		if algorithm == rowDifference {
			sum += checksumRowDifference(str)
		}
		if algorithm == rowDivisible {
			sum += checksumRowDivisible(str)
		}
	}

	return sum
}

func checksumRowDifference(input string) int {
	var largest = 0
	var smallest = 65536

	for _, str := range strings.Split(input, "\t") {
		if str == "" {
			// if we couldn't parse a value, skip to next iteration of loop
			continue
		}

		i, _ := strconv.Atoi(str)
		if i > largest {
			largest = i
		}
		if i < smallest {
			smallest = i
		}
	}

	return largest - smallest
}

func checksumRowDivisible(input string) int {
	// load row's values into an integer array
	rowValues := strings.Split(input, "\t")
	var values = make([]int, len(rowValues))
	for i, val := range rowValues {
		values[i], _ = strconv.Atoi(val)
	}

	// sort the values
	sort.Ints(values)

	// iterate from both directions of the array to find the evenly divisible pair
	var quotient = 0
	for j := len(values) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if values[i] == 0 { continue }  // skip on what was probably a bad parse
			if values[j] % values[i] == 0 {
				quotient = values[j] / values[i]
				return quotient
			}
		}
	}

	return 0  // throw error?
}