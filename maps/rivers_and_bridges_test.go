package maps_test

import (
	"testing"

	"github.com/BattlesnakeOfficial/rules"
	"github.com/BattlesnakeOfficial/rules/maps"
	"github.com/stretchr/testify/require"
)

func TestRiversAndBridgetsHazardsMap(t *testing.T) {
	// check error handling
	m := maps.RiverAndBridgesMediumHazardsMap{}
	settings := rules.Settings{}

	// check error for unsupported board sizes
	state := rules.NewBoardState(9, 9)
	editor := maps.NewBoardStateEditor(state)
	err := m.SetupBoard(state, settings, editor)
	require.Error(t, err)

	tests := []struct {
		Map    maps.GameMap
		Width  uint
		Height uint
	}{
		{maps.RiverAndBridgesMediumHazardsMap{}, 11, 11},
		{maps.RiverAndBridgesLargeHazardsMap{}, 19, 19},
		{maps.RiverAndBridgesExtraLargeHazardsMap{}, 25, 25},
	}

	// check all the supported sizes
	for _, test := range tests {
		state = rules.NewBoardState(int(test.Width), int(test.Height))
		state.Snakes = append(state.Snakes, rules.Snake{ID: "1", Body: []rules.Point{}})
		editor = maps.NewBoardStateEditor(state)
		require.Empty(t, state.Hazards)
		err = test.Map.SetupBoard(state, settings, editor)
		require.NoError(t, err)
		require.NotEmpty(t, state.Hazards)
	}
}
