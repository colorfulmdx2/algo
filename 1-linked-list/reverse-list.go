package linked_list

func reverseList(head *ListNode) *ListNode {
	var newNext *ListNode // нода на которую мы будем переписовать связь
	for head != nil {
		currentNext := head.Next // делаем копию значения, так как мы перезапишем его ниже
		head.Next = newNext      // переписываем связь
		newNext = head           // переписываем ноду для следующей связи на + 1
		head = currentNext       // двигаем цикл

		// решение с паралельным присваиванием
		/*Присваивание происходит в две фазы. Сначала вычисляются операнды выражений
		индексов и косвенных указателей (включая неявные косвенные указатели в селекторах)
		слева и выражения справа в обычном порядке. Затем присваивания выполняются слева
		направо.*/

		/*head.Next, newNext, head = newNext, head, head.Next*/
	}
	return newNext
}
