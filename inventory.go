package hotomata

import (
	"encoding/json"
)

type InventoryMachine struct {
	Name       string
	Groups     InventoryGroups
	properties map[string]json.RawMessage
}

type InventoryGroup struct {
	GroupName  string
	properties map[string]json.RawMessage
}

type InventoryGroups []InventoryGroup

func (groups InventoryGroups) Names() (names []string) {
	for _, g := range groups {
		names = append(names, g.GroupName)
	}
	return
}

func ParseInventory(inventoryJson []byte) ([]InventoryMachine, error) {
	var items []map[string]json.RawMessage
	err := json.Unmarshal(inventoryJson, &items)
	if err != nil {
		return []InventoryMachine{}, err
	}

	inventoryMachines, err := parseInventoryItems(InventoryGroups{}, items)
	return inventoryMachines, err
}

func parseInventoryItems(groups InventoryGroups, items []map[string]json.RawMessage) ([]InventoryMachine, error) {
	var inventoryMachines = []InventoryMachine{}

	for _, item := range items {
		// try handling item as a machine
		var machineName string
		if err := json.Unmarshal(item["name"], &machineName); err == nil && machineName != "" {
			delete(item, "name")
			inventoryMachines = append(inventoryMachines, InventoryMachine{
				Name:       machineName,
				Groups:     groups,
				properties: item,
			})
		}

		// try handling item as a group
		var groupName string
		if err := json.Unmarshal(item["group_name"], &groupName); err == nil && groupName != "" {
			var groupItems []map[string]json.RawMessage
			err := json.Unmarshal(item["machines"], &groupItems)
			if err != nil {
				return inventoryMachines, err
			}

			delete(item, "group_name")
			delete(item, "machines")

			group := InventoryGroup{
				GroupName:  groupName,
				properties: item,
			}
			newItems, err := parseInventoryItems(append(groups, group), groupItems)
			if err != nil {
				return inventoryMachines, err
			}
			inventoryMachines = append(inventoryMachines, newItems...)
		}
	}

	return inventoryMachines, nil
}
