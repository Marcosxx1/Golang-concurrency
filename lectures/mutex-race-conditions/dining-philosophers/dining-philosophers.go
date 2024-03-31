package diningphilosophers

/*

Rules:
1 - One fork per philosopher: Each philosopher needs two forks to eat a meal.

2 - Alternating actions: Philosophers alternate between thinking and eating. While thinking, they do not interact
 with forks. While eating, they must hold two forks simultaneously.

3 - Deadlock prevention: To prevent deadlock, philosophers must follow certain rules when acquiring forks. Deadlock
 occurs when each philosopher holds one fork and is waiting for another, creating a circular dependency.

4 - Mutual exclusion: Only one philosopher can hold a fork at any given time. This ensures that philosophers do not
 try to eat from the same fork simultaneously.

5 - Starvation prevention: Every philosopher should have an opportunity to eat eventually. Starvation occurs when a
 philosopher is unable to acquire both forks indefinitely due to the actions of other philosophers.


	A nice approach could be to allow each philosopher to pick up only the left fork first
and then try to pick up the right fork. If the right fork is not available, they put down
the left fork and wait for a while before trying again. This way, you can maximize resource
utilization by allowing multiple philosophers to eat simultaneously if possible.

Here's a rough outline of how we will (try to) implement this:

Create a slice of forks, where each philosopher corresponds to a fork.
Each philosopher tries to pick up the left fork first.
If successful, they try to pick up the right fork.
If both forks are available, they eat for a while.
After eating, they release both forks and start thinking again.
If the right fork is not available, they release the left fork and wait for a random duration before trying again. */

var philosophers = []string{"Plato", "Aristotle", "Socrates", "Confucius", "Immanuel Kant"}
