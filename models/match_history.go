package models

import (
	"time"
)
// MatchHistory represents match history
type MatchHistory struct {
    BaseModel
    Name          string    `json:"name"`
    TeamID        string       `json:"team_id"`
    Opponent      string    `json:"opponent"`
    MatchDate     time.Time `json:"match_date"`
    ScoreTeam     int       `json:"score_team"`
    ScoreOpponent int       `json:"score_opponent"`
    Description   string    `json:"description,omitempty"`
}
