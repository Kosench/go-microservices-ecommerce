package part

import (
	"context"
	"errors"
	"sync"

	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/model"
	"github.com/Kosench/go-microservices-ecommerce/inventory/internal/repository/converter"
	repoModel "github.com/Kosench/go-microservices-ecommerce/inventory/internal/repository/model"
)

var (
	ErrNotFound      = errors.New("part not found")
	ErrAlreadyExists = errors.New("part already exists")
)

type memoryInventoryRepository struct {
	mu   sync.RWMutex
	data map[string]*repoModel.Part
}

func NewMemoryInventoryRepository() *memoryInventoryRepository {
	return &memoryInventoryRepository{
		data: make(map[string]*repoModel.Part),
	}
}

// GetPart принимает uuid, возвращает model.Part
func (r *memoryInventoryRepository) GetPart(ctx context.Context, uuid string) (*model.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// 1. Найти в "БД" (map) → получаем repoModel.Part
	repoPart, exists := r.data[uuid]
	if !exists {
		return nil, ErrNotFound
	}

	// 2. Конвертировать repoModel.Part → model.Part
	servicePart := converter.ConvertRepoPartToServicePart(repoPart)

	return servicePart, nil
}

func (r *memoryInventoryRepository) ListParts(ctx context.Context, filter *model.PartsFilter) ([]*model.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// 1. Конвертировать service filter → repo filter
	repoFilter := converter.ConvertServiceFilterToRepoFilter(filter)

	// 2. Собрать все части из "БД"
	allRepoParts := make([]*repoModel.Part, 0, len(r.data))
	for _, repoPart := range r.data {
		allRepoParts = append(allRepoParts, repoPart)
	}

	// 3. Фильтрация (работаем с repoModel.Part)
	filteredRepoParts := r.filterParts(allRepoParts, repoFilter)

	// 4. Конвертировать []repoModel.Part → []*model.Part
	serviceParts := make([]*model.Part, len(filteredRepoParts))
	for i, repoPart := range filteredRepoParts {
		serviceParts[i] = converter.ConvertRepoPartToServicePart(repoPart)
	}

	return serviceParts, nil
}

// filterParts фильтрует repoModel.Part по repoModel.PartsFilter
// Логика: ИЛИ внутри поля, И между полями
func (r *memoryInventoryRepository) filterParts(parts []*repoModel.Part, filter *repoModel.PartsFilter) []*repoModel.Part {
	if filter == nil || r.isFilterEmpty(filter) {
		return parts
	}

	result := parts

	// Шаг 1: Фильтр по UUID (ИЛИ)
	if len(filter.Uuids) > 0 {
		filtered := make([]*repoModel.Part, 0)
		for _, part := range result {
			if r.containsString(filter.Uuids, part.UUID) {
				filtered = append(filtered, part)
			}
		}
		result = filtered
	}

	// Шаг 2: Фильтр по имени (ИЛИ)
	if len(filter.Names) > 0 {
		filtered := make([]*repoModel.Part, 0)
		for _, part := range result {
			if r.containsString(filter.Names, part.Name) {
				filtered = append(filtered, part)
			}
		}
		result = filtered
	}

	// Шаг 3: Фильтр по категории (ИЛИ)
	if len(filter.Categories) > 0 {
		filtered := make([]*repoModel.Part, 0)
		for _, part := range result {
			if r.containsCategory(filter.Categories, part.Category) {
				filtered = append(filtered, part)
			}
		}
		result = filtered
	}

	// Шаг 4: Фильтр по стране производителя (ИЛИ)
	if len(filter.ManufacturerCountries) > 0 {
		filtered := make([]*repoModel.Part, 0)
		for _, part := range result {
			if part.Manufacturer != nil && r.containsString(filter.ManufacturerCountries, part.Manufacturer.Country) {
				filtered = append(filtered, part)
			}
		}
		result = filtered
	}

	// Шаг 5: Фильтр по тегам (ИЛИ) - хотя бы один тег совпадает
	if len(filter.Tags) > 0 {
		filtered := make([]*repoModel.Part, 0)
		for _, part := range result {
			if r.hasAnyTag(part.Tags, filter.Tags) {
				filtered = append(filtered, part)
			}
		}
		result = filtered
	}

	return result
}

// Вспомогательные методы для фильтрации
func (r *memoryInventoryRepository) isFilterEmpty(filter *repoModel.PartsFilter) bool {
	return len(filter.Uuids) == 0 &&
		len(filter.Names) == 0 &&
		len(filter.Categories) == 0 &&
		len(filter.ManufacturerCountries) == 0 &&
		len(filter.Tags) == 0
}

func (r *memoryInventoryRepository) containsString(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func (r *memoryInventoryRepository) containsCategory(slice []repoModel.Category, value repoModel.Category) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func (r *memoryInventoryRepository) hasAnyTag(partTags []string, filterTags []string) bool {
	for _, partTag := range partTags {
		for _, filterTag := range filterTags {
			if partTag == filterTag {
				return true
			}
		}
	}
	return false
}
