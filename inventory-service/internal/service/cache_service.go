package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

type CacheService struct {
	productService *productService // Use concrete type instead of interface
	cacheStore     *cache.Cache
}

func NewCacheService(ps *productService, cs *cache.Cache) *CacheService {
	return &CacheService{
		productService: ps,
		cacheStore:     cs,
	}
}
func (s *CacheService) StartCacheRefresh(ctx context.Context) {
	ticker := time.NewTicker(12 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.RefreshCache(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (s *CacheService) RefreshCache(ctx context.Context) error {
	// Clear all cache
	s.cacheStore.Flush()
	log.Println("Cache flushed successfully")

	// Warm up cache by fetching all products
	products, err := s.productService.GetAllProducts(ctx)
	if err != nil {
		log.Printf("Failed to refresh cache: %v", err)
		return fmt.Errorf("failed to refresh cache: %w", err)
	}

	log.Printf("Cache refreshed successfully with %d products", len(products))
	return nil
}
