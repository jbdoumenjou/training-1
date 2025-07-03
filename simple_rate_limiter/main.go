package main

//  📝 Problem Statement
//  Build a simple in-memory rate limiter that can:
//  Limit the number of requests per user (by user ID)
//  Provide a function Allow(userID string) bool that returns whether a request should be allowed
//  Support bursting (i.e., a user can send up to N requests at once)
//  Support a rate of M requests per second
//  Write unit tests to verify the behavior.

// 🧱 Requirements
//  * Support configurable rate and burst size per user
//  * Thread-safe (Allow() may be called concurrently)
//  * Can use Go standard library only (or testify for tests)
//  * Good separation (struct, method, possibly interface)
//  * Bonus: expiry of inactive users (optional cleanup)

type RateLimiter interface {
	Allow(userID string) bool
}

func main() {

}
