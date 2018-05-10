package Decorator

import (
	"errors"
)

type IngredientAdd interface {
	AddIngredient() (string, error)
}

/*Decorator*/
type PizzaDecorator struct {
	Ingredient IngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

/* Meat Ingredient*/
type Meat struct {
	Ingredient IngredientAdd
}

func (m *Meat) AddIngredient() (string, error) {
	return "", errors.New("Not implemented yet")
}

/* Meat Ingredient*/
type Onion struct {
	Ingredient IngredientAdd
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Onion")
	}

	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s,", s, "onion"), nil
}

func (m *Meat) AddIngredient() (string, error) {

	if m.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Meat")
	}

	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s,", s, "meat"), nil
}
