package main

import (
	"math"
	"os"
	"io"
	"bufio"
	"strings"
	"strconv"
	"fmt"
	"sort"
)

type JunctionBox struct {
	x float64
	y float64
	z float64
	nextBox *JunctionBox
	circut *Circut
}

type JunctionBoxConnection struct {
	from *JunctionBox
	to *JunctionBox
	dist float64
}

type Circut struct {
	length int
	start *JunctionBox
}

func makeCircut(head *JunctionBox) *Circut {
	output := &Circut{
		length: 1,
		start: head,
	}
	head.circut = output;
	return output;
}

func addJuctionToCircut(c *Circut, j *JunctionBox) {
	j.circut = c
	if c.start == nil {
		c.start = j;
	} else {

		n := c.start
		for n.nextBox != nil {
			n = n.nextBox;
		}

		n.nextBox = j;
	}

	c.length++;
}

func mergeCircut(a, b *Circut) {

	n := a.start
	for n.nextBox != nil {
		n = n.nextBox;
		n.circut = a;
	}

	n.nextBox = b.start;
	
	m := b.start;
	for m != nil {
		m.circut = a;
		m = m.nextBox;
	}

	a.length += b.length;
}

func makeJunctionBox(x, y, z float64) *JunctionBox {
	return &JunctionBox{
		x: x,
		y: y,
		z: z,
		nextBox: nil,
		circut: nil,
	};
}

func makeJunctionBoxFromString(str string) *JunctionBox {
	str = strings.TrimSpace(str)
	strSplit := strings.Split(str, ",");
	x, _ := strconv.Atoi(strSplit[0]);
	y, _ := strconv.Atoi(strSplit[1]);
	z, _ := strconv.Atoi(strSplit[2]);

	return makeJunctionBox(float64(x), float64(y), float64(z));
}

func getDistance(a, b *JunctionBox) float64 {
    dx := a.x - b.x
    dy := a.y - b.y
    dz := a.z - b.z
    return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func getCloserJuctionBox(a, b, c *JunctionBox) *JunctionBox {
    if getDistance(a, b) <= getDistance(a, c) {
        return b
    }
    return c
}

func makeKShortestConnections(list *[]*JunctionBox, k int) {

    junctionBoxes := *list;
    var allConnections []JunctionBoxConnection;

    for i := 0; i < len(junctionBoxes); i++ {
        for j := i + 1; j < len(junctionBoxes); j++ {
            a := junctionBoxes[i];
            b := junctionBoxes[j];
            dist := getDistance(a, b);
            
            allConnections = append(allConnections, JunctionBoxConnection{
                from: a,
                to: b,
                dist: dist,
            })
        }
    }

    sort.Slice(allConnections, func(i, j int) bool {
        return allConnections[i].dist < allConnections[j].dist
    })

    for i := 0; i < k && i < len(allConnections); i++ {
        conn := allConnections[i]
        boxA := conn.from
        boxB := conn.to

        if boxA.circut == nil {
            makeCircut(boxA)
        }
        if boxB.circut == nil {
            makeCircut(boxB)
        }

        if boxA.circut != boxB.circut {
            mergeCircut(boxA.circut, boxB.circut)
        }
    }
}

func printCircut(c *Circut) {
	fmt.Printf("Circut -> ");

	n := c.start;
	for n != nil {
		fmt.Printf("Juction[%.0f, %.0f, %.0f] -> ", n.x, n.y, n.z);
		if (n.nextBox == n) {
			fmt.Println("Cycle!");
			return;
		}
		n = n.nextBox;
	}
	fmt.Println(" nil");
}

func printBoxJunctions(junctionBoxes []*JunctionBox) {
	for i := 0; i < len(junctionBoxes); i++ {
		fmt.Println(*junctionBoxes[i])
		if (*junctionBoxes[i]).circut != nil {
			printCircut((*junctionBoxes[i]).circut);
		}
	}
}

func readFile(name string) (string, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	str := string(b)
    return str, nil
}

func main() {

	fileStr, _ := readFile("question8.input");
	reader := bufio.NewReader(strings.NewReader(fileStr));
	junctionBoxes := []*JunctionBox{};
	connections := 1000;

	for {
		line, err := reader.ReadString('\n');
		if err == io.EOF {
			break;
		} else {
			j := makeJunctionBoxFromString(line);
			junctionBoxes = append(junctionBoxes, j);
		}
	}

	makeKShortestConnections(&junctionBoxes, connections);

	circutCount := make(map[*Circut]int);

	for _, box := range junctionBoxes {
		if box.circut != nil {
			circutCount[box.circut] = box.circut.length;
		}
	}

	getLargest := 3
	total := 1
	for i := 0; i < getLargest; i++ {
		
		var largestCircut *Circut = nil;
		largestLength := 1;
		for c, length := range circutCount {
			if (largestLength < length) {
				largestCircut = c;
				largestLength = length;
			}
		}
		fmt.Println(largestLength);
		total *= largestLength;
		circutCount[largestCircut] = 0;
	}

	fmt.Println(total);
}