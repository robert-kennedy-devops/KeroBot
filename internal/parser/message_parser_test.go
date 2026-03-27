package parser

import "testing"

func TestParseStatePrecedence(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		text    string
		buttons []string
		want    GameState
	}{
		{
			name:    "victory_over_hunting",
			text:    "Você ganhou! Vitória! Continue caçando.",
			buttons: []string{"Caçar"},
			want:    StateVictory,
		},
		{
			name:    "combat_over_menu",
			text:    "Um inimigo apareceu.",
			buttons: []string{"Caçar", "Atacar"},
			want:    StateCombat,
		},
		{
			name:    "defeat_over_menu",
			text:    "Derrota... tente novamente.",
			buttons: []string{"Caçar"},
			want:    StateDefeat,
		},
		{
			name:    "inventory_when_button_or_text",
			text:    "Abrindo inventario",
			buttons: []string{"Inventário"},
			want:    StateInventory,
		},
		{
			name:    "dungeon_when_button_present",
			text:    "Prepare-se para a masmorra",
			buttons: []string{"Masmorra"},
			want:    StateDungeon,
		},
		{
			name:    "hunting_when_text_only",
			text:    "Você está caçando na floresta",
			buttons: []string{},
			want:    StateHunting,
		},
		{
			name:    "main_menu_when_hunt_button_only",
			text:    "Menu principal",
			buttons: []string{"Caçar"},
			want:    StateMainMenu,
		},
		{
			name:    "main_menu_when_inventory_button_present",
			text:    "Menu principal",
			buttons: []string{"⚔️ Caçar", "🧳 Inventário", "🗺️ Viajar"},
			want:    StateMainMenu,
		},
		{
			name:    "main_menu_when_masmorra_keys_info",
			text:    "Chaves de Masmorra: 0",
			buttons: []string{"⚔️ Caçar", "🗝️ Masmorra"},
			want:    StateMainMenu,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Parse(tt.text, tt.buttons)
			if got.State != tt.want {
				t.Fatalf("state=%s want=%s", got.State, tt.want)
			}
		})
	}
}

func TestParseHPAndPotions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		text      string
		wantHP    int
		wantPots  int
		wantEnergy int
		wantEnergyMax int
	}{
		{
			name:     "hp_percent_simple",
			text:     "HP: 35%",
			wantHP:   35,
			wantPots: -1,
		},
		{
			name:     "hp_fraction",
			text:     "HP 120/300",
			wantHP:   40,
			wantPots: -1,
		},
		{
			name:     "potions_poçoes",
			text:     "Poções: 7",
			wantHP:   0,
			wantPots: 7,
		},
		{
			name:     "potions_estoque",
			text:     "Estoque: 3",
			wantHP:   0,
			wantPots: 3,
		},
		{
			name:     "potions_item_count",
			text:     "Poção de Vida x12",
			wantHP:   0,
			wantPots: 12,
			wantEnergy: 0,
			wantEnergyMax: 0,
		},
		{
			name:     "energy_fraction",
			text:     "⚡ Energia: 3/20",
			wantHP:   0,
			wantPots: -1,
			wantEnergy: 3,
			wantEnergyMax: 20,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Parse(tt.text, nil)
			if got.HPPercent != tt.wantHP {
				t.Fatalf("hp=%d want=%d", got.HPPercent, tt.wantHP)
			}
			if got.Potions != tt.wantPots {
				t.Fatalf("potions=%d want=%d", got.Potions, tt.wantPots)
			}
			if got.Energy != tt.wantEnergy || got.EnergyMax != tt.wantEnergyMax {
				t.Fatalf("energy=%d/%d want=%d/%d", got.Energy, got.EnergyMax, tt.wantEnergy, tt.wantEnergyMax)
			}
		})
	}
}

func TestParseNoEnergyState(t *testing.T) {
	t.Parallel()

	got := Parse("Sem energia para caçar.", []string{"⚡ Energia", "🏠 Menu"})
	if got.State != StateNoEnergy {
		t.Fatalf("state=%s want=%s", got.State, StateNoEnergy)
	}
}
