package parser

type GameState string

const (
	StateUnknown   GameState = "UNKNOWN"
	StateMainMenu  GameState = "MAIN_MENU"
	StateHunting   GameState = "HUNTING"
	StateCombat    GameState = "COMBAT"
	StateVictory   GameState = "VICTORY"
	StateDefeat    GameState = "DEFEAT"
	StateInventory GameState = "INVENTORY"
	StateDungeon   GameState = "DUNGEON"
	StateNoEnergy  GameState = "NO_ENERGY"
)

type Snapshot struct {
	State        GameState
	HPPercent    int
	Potions      int
	Energy       int
	EnergyMax    int
	Buttons      []string
	Text         string
}
