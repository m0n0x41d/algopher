// For testing ease and class completeness, additional methods are implemented for the same struct in the task2.go file,
// since it does not violate laboratory server rules – it is allowed to add methods in the type signature.

// 9.* Add a method that reverses the order of elements in a linked list.

// 10.* Add a Boolean method that determines whether there are cycles (loop-like loops) within the list.

// 11.* Add a method that sorts the list.

// 12.* Add a method that merges two lists into a third. The values specified by these prerequisites must be sorted and the resulting list must be returned with all elements also in order (a sorting method is not allowed for the resulting list).

// Additional reflexes:
// As I mentioned in the comment on the Merge method, the linked list interface design continues to reveal
// its shortcomings - it requires the use of structure fields to implement "external methods."
// Even though these methods are external exclusively to the user, it seems they should be "external"
// primarily to the implementation itself.
// 1. About violating encapsulation
// The Merge method directly manipulates left.head and right.head,
// bypassing the public interface. This violates encapsulation and makes
// the code brittle - changing the internal structure of LinkedList2
// will break Merge, even though it is formally an "external" method.
//
// 2. On the need for the Pop/Shift operation
// The absence of the PopHead() method (remove and return the head) forces us to
// write `left.head = left.head.next` directly, which:
// 1) Doesn't update the length
// 2) Doesn't clear the prev of the new head
// 3) Doesn't handle the case when the list becomes empty
// A proper interface should provide atomic operations.
//
// 3. On the receiver of the Merge method
// Merge is called as l.Merge(left, right), but the receiver `l`
// is not used at all—the method creates a new list.
// This should be either a function (not a method) or a constructor:
// func MergeLists(left, right *LinkedList2) *LinkedList2
//
// 4. On the mutation of input parameters
// Merge mutates the input lists (left.head, right.head change).
// After calling Merge, the original lists are in an invalid
// state — head is shifted, but length is not updated.
// Ideally, everything should be rewritten in the spirit of functional immutability
// and return new list instances, albeit at the cost of memory – Go's GC is good,
// and the additional guarantees provided by immutability are invaluable
// Ideally, everything should be rewritten in the spirit of functional immutability
// and return new list instances, albeit at the cost of memory – Go's GC is good,
// and the additional guarantees provided by immutability are invaluable.
//
// 5. General Conclusion
// Conclusion: A linked list as a data structure is conceptually simple,
// but designing a clean API for it is non-trivial (as for any non-trivial type).
// Each new operation reveals the flaws of the underlying interface.
// It would certainly be more correct to start by defining the full set of operations
// via interfaces, following the Abstract Data Types approach. But we know that's the context of a separate course.
