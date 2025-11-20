package main

import (
	"testing"
)

// TestRSSFeedWithInvalidURL tests that rssFeed returns an empty slice
// instead of panicking when given an invalid URL
func TestRSSFeedWithInvalidURL(t *testing.T) {
	// Test with a completely invalid URL
	result := rssFeed("invalid-url", 5)
	
	if result == nil {
		t.Error("Expected non-nil slice, got nil")
	}
	
	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %d items", len(result))
	}
}

// TestRSSFeedWithNonExistentDomain tests that rssFeed returns an empty slice
// when given a URL that doesn't exist
func TestRSSFeedWithNonExistentDomain(t *testing.T) {
	// Test with a non-existent domain
	result := rssFeed("https://this-domain-does-not-exist-12345.com/feed.xml", 5)
	
	if result == nil {
		t.Error("Expected non-nil slice, got nil")
	}
	
	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %d items", len(result))
	}
}

// TestRSSFeedWithInvalidFeed tests that rssFeed returns an empty slice
// when the URL doesn't point to a valid RSS feed
func TestRSSFeedWithInvalidFeed(t *testing.T) {
	// Test with a URL that exists but isn't an RSS feed
	result := rssFeed("https://www.google.com", 5)
	
	if result == nil {
		t.Error("Expected non-nil slice, got nil")
	}
	
	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %d items", len(result))
	}
}
