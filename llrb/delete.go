package llrb

import "fmt"

// DeleteMin deletes the minimum-Key Node of the sub-Tree.
func DeleteMin(nd *Node) (*Node, Interface) {
	fmt.Println("DeleteMin", nd.Key)
	if nd == nil {
		return nil, nil
	}
	if nd.Left == nil {
		return nil, nd.Key
	}
	if !isRed(nd.Left) && !isRed(nd.Left.Left) {
		nd = MoveRedFromRightToLeft(nd)
	}
	var deleted Interface
	nd.Left, deleted = DeleteMin(nd.Left)
	return FixUp(nd), deleted
}

// DeleteMin deletes the minimum-Key Node of the Tree.
// It returns the minimum Key or nil.
func (tr *Tree) DeleteMin() Interface {
	var deleted Interface
	tr.Root, deleted = DeleteMin(tr.Root)
	if tr.Root != nil {
		tr.Root.Black = true
	}
	return deleted
}

// DeleteMax deletes the maximum-Key Node of the sub-Tree.
func DeleteMax(nd *Node) (*Node, Interface) {
	if nd == nil {
		return nil, nil
	}
	if nd.Left == nil {
		return nil, nd.Key
	}
	if !isRed(nd.Right) && !isRed(nd.Right.Left) {
		nd = MoveRedFromRightToLeft(nd)
	}
	var deleted Interface
	nd.Right, deleted = DeleteMax(nd.Right)
	return FixUp(nd), deleted
}

// DeleteMax deletes the maximum-Key Node of the Tree.
// It returns the maximum Key or nil.
func (tr *Tree) DeleteMax() Interface {
	var deleted Interface
	tr.Root, deleted = DeleteMax(tr.Root)
	if tr.Root != nil {
		tr.Root.Black = true
	}
	return deleted
}

// Delete deletes the node with the Key and returns the Key Interface.
// It returns nil if the Key does not exist in the tree.
//
//
//	Delete Algorithm:
//	1. Start 'delete' from tree Root.
//
//	2. Call 'delete' method recursively on each Node from binary search path.
//		- e.g. If the key to delete is greater than Root's key
//			, call 'delete' on Right Node.
//
//
//	# start
//
//	3. Recursive 'tree.delete(nd, key)'
//
//		if key < nd.Key:
//
//			if nd.Left == empty:
//				print "then nothing to delete, so return nil"
//				return nd, nil
//
//			if (nd.Left == Black) and (nd.Left.Left == Black):
//				print "then move Red from Right to Left to update nd"
//				nd = MoveRedFromRightToLeft(nd)
//
//			print "recursively call 'delete' to update nd.Left"
//			nd.Left, deleted = tr.delete(nd.Left, key)
//
//		else if key >= nd.Key:
//
//			if nd.Left == Red:
//				print "then RotateToRight(nd) to update nd"
//				nd = RotateToRight(nd)
//
//			if (key == nd.Key) and nd.Right == empty:
//				print "then return nil, nd.Key to recursively update nd"
//				return nil, nd.Key
//
//			if (nd.Right != empty)
//			and (nd.Right == Black)
//			and (nd.Right.Left == Black):
//				print "then move Red from Left to Right to update nd"
//				nd = MoveRedFromLeftToRight(nd)
//
//			# nd gets updated by this step
//
//			if (key == nd.Key):
//				print "then DeleteMin of nd.Right to update nd.Right"
//				nd.Right, subDeleted = DeleteMin(nd.Right)
//
//				print "and then update nd.Key with DeleteMin(nd.Right)"
//				deleted, nd.Key = nd.Key, subDeleted
//
//			else if key != nd.Key:
//				print "recursively call 'delete' to update nd.Right"
//				nd.Right, deleted = tr.delete(nd.Right, key)
//
//	# end
//
//
//	4. If the tree's Root is not nil, set Root Black.
//
//	5. Return the Interface(nil if the key does not exist.)
//
func (tr *Tree) Delete(key Interface) Interface {
	var deleted Interface
	tr.Root, deleted = tr.delete(tr.Root, key)
	if tr.Root != nil {
		tr.Root.Black = true
	}
	return deleted
}

func (tr *Tree) delete(nd *Node, key Interface) (*Node, Interface) {
	fmt.Println("calling delete on", nd.Key, "for the key", key)
	if nd == nil {
		return nil, nil
	}

	var deleted Interface

	// if key is Less than nd.Key
	if key.Less(nd.Key) {

		// if key is Less than nd.Key
		// and nd.Left is empty
		if nd.Left == nil {

			// then nothing to delete
			// so return the nil
			return nd, nil
		}

		// if key is Less than nd.Key
		// and nd.Left is Black
		// and nd.Left.Left is Black
		if !isRed(nd.Left) && !isRed(nd.Left.Left) {

			// then MoveRedFromRightToLeft(nd)
			nd = MoveRedFromRightToLeft(nd)
		}

		// and recursively call tr.delete(nd.Left, key)
		nd.Left, deleted = tr.delete(nd.Left, key)

	} else {
		// if key is not Less than nd.Key
		//(or key is greater than or equal to nd.Key)
		//(or key >= nd.Key)

		// and nd.Left is Red
		if isRed(nd.Left) {

			// then RotateToRight(nd)
			nd = RotateToRight(nd)
			fmt.Println("after nd = RotateToRight(nd)", nd.Key)
		}

		// and nd.Key is not Less than key
		// (or nd.Key >= key)
		// (or key == nd.Key)
		// and nd.Right is empty
		if !nd.Key.Less(key) && nd.Right == nil {
			fmt.Println("!nd.Key.Less(key) && nd.Right == nil when", nd.Key)
			// then return nil to delete the key
			return nil, nd.Key
		}

		// and nd.Right is not empty
		// and nd.Right is Black
		// and nd.Right.Left is Black
		if nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) {
			fmt.Println("nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when", nd.Key)

			// then MoveRedFromLeftToRight(nd)
			nd = MoveRedFromLeftToRight(nd)
		}

		// and key == nd.Key
		if !nd.Key.Less(key) {

			var subDeleted Interface

			// then DeleteMin(nd.Right)
			nd.Right, subDeleted = DeleteMin(nd.Right)
			if subDeleted == nil {
				panic("Unexpected nil value")
			}

			// and update nd.Key with DeleteMin(nd.Right)
			deleted, nd.Key = nd.Key, subDeleted

		} else {
			// if updated nd.Key is Less than key (nd.Key < key) to update nd.Right
			fmt.Println("nd.Right, deleted = tr.delete(nd.Right, key) at", nd.Key)
			nd.Right, deleted = tr.delete(nd.Right, key)
		}
	}

	return FixUp(nd), deleted
}
