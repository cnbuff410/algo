/*
* Disjoint Sets. Refer to
* http://www.mathblog.dk/disjoint-set-data-structure/?utm_source=feedburner&utm_medium=feed&utm_campaign=Feed%3A+mathblogdk+%28MathBlog%29
* 
* Author: likunarmstrong@gmail.com
 */

package disjoint

type DisJointSet struct {
	parent     []int // Whole universe
	rank       []int // The rank of each element in the universe
	totalCount int   // Total number of elements in this set
	setCount   int   // Total number of sets
	sizeOfSet  []int // Size of each set
}

// Initializes a new Disjoint-Set data structure, with the specified amount of
// elements in the universe.
func NewSet(count int) *DisJointSet {
	parent := make([]int, count)
	rank := make([]int, count)
	totalCount := count
	setCount := count
	sizeOfSet := make([]int, count)
	for i := 0; i < count; i++ {
		parent[i] = i
		rank[i] = 0
		sizeOfSet[i] = 1
	}

	s := &DisJointSet{parent, rank, totalCount, setCount, sizeOfSet}
	return s
}

// Find the parent of the specified element.
func (s *DisJointSet) Find(i int) int {
	if s.parent[i] == i {
		return i
	}
	// Path compression
	// i is not parent of itself. Recursively find the real parent of i,
	// then cache it by moving i's node directly under the representative
	// of its set
	s.parent[i] = s.Find(s.parent[i])
	return s.parent[i]
}

// Unite the sets that the specified elements belong to, by rank
func (s *DisJointSet) Union(i, j int) {
	// Root node of i, j
	iRoot := s.Find(i)
	jRoot := s.Find(j)
	// Elements are in the same set, no need to unite anything.
	if iRoot == jRoot {
		return
	}

	// Rank of i, j's tree
        iRank := s.rank[iRoot]
        jRank := s.rank[jRoot]

	s.setCount--

	// If i's rank is less than j's rank
	if iRank < jRank {
		// Then move i under j
		s.parent[iRoot] = jRoot
		s.sizeOfSet[jRoot] += s.sizeOfSet[iRoot]
	} else if jRank < iRank {
		// Then move j under i
		s.parent[jRoot] = iRoot
		s.sizeOfSet[iRoot] += s.sizeOfSet[jRoot]
	} else { // Same rank
		// Then move i under j (doesn't matter which one goes where)
		s.parent[iRoot] = jRoot
		s.sizeOfSet[jRoot] += s.sizeOfSet[iRoot]
		// And increment the the result tree's rank by 1
		s.rank[iRoot]++
	}
}

// Return the number of elements in the set that the specified elements belong to.
func (s *DisJointSet) SetSize(i int) int {
	return s.sizeOfSet[s.Find(i)]
}
