package preview_test

import (
	"testing"

	"botfinac/tests/internal/contract"
)

func TestPreviewGlobalMarketScenarios(t *testing.T) {
	document, err := contract.ReadArchitecture()
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []contract.Scenario{
		{Name: "sandbox_equity_loss_limit", Context: "US equities; Phase 0 sandbox; loss-limit tripwire; reject order", Rule: "cada limite em sandbox"},
		{Name: "sandbox_daily_trade_cap", Context: "US index futures; Phase 0; active session; max-trade cap", Rule: "cada limite em sandbox"},
		{Name: "sandbox_rr_floor", Context: "FX majors; Phase 0; London session; minimum risk-reward gate", Rule: "cada limite em sandbox"},
		{Name: "sandbox_cooldown", Context: "Brazil equities; Phase 0; post-loss session; cooldown gate", Rule: "cada limite em sandbox"},
		{Name: "sandbox_market_hours", Context: "Tokyo equities; Phase 0; opening auction; schedule boundary", Rule: "cada limite em sandbox"},
		{Name: "sandbox_state_veto", Context: "EuroStoxx futures; Phase 0; volatile open; state veto", Rule: "cada limite em sandbox"},
		{Name: "sandbox_oco_coverage", Context: "Crude futures; Phase 0; overnight; broker OCO primitive", Rule: "cada limite em sandbox"},
		{Name: "sandbox_reconnect_fill", Context: "Gold futures; Phase 0; event spike; reconnect after partial fill", Rule: "cada limite em sandbox"},
		{Name: "sandbox_duplicate_order", Context: "USD/JPY; Phase 0; Tokyo session; idempotent resend", Rule: "cada limite em sandbox"},
		{Name: "sandbox_guardian_stop", Context: "Crypto derivatives; Phase 0; weekend; independent guardian stop", Rule: "cada limite em sandbox"},
		{Name: "demo_signal_count_us", Context: "US equity index; Phase 1; cash session; demo signal count gate", Rule: "≥ 200 sinais gravados"},
		{Name: "demo_signal_count_eu", Context: "European equities; Phase 1; open; demo signal count gate", Rule: "≥ 200 sinais gravados"},
		{Name: "demo_signal_count_asia", Context: "Asia equity futures; Phase 1; overnight; demo signal count gate", Rule: "≥ 200 sinais gravados"},
		{Name: "demo_signal_count_fx", Context: "G10 FX; Phase 1; overlapping sessions; demo signal count gate", Rule: "≥ 200 sinais gravados"},
		{Name: "demo_signal_count_brl", Context: "BRL FX; Phase 1; local policy session; demo signal count gate", Rule: "≥ 200 sinais gravados"},
		{Name: "demo_spread_cost_us", Context: "US equities; Phase 1; opening gap; spread-adjusted expectancy", Rule: "expectativa positiva DEPOIS de custos"},
		{Name: "demo_fee_cost_eu", Context: "European index futures; Phase 1; regular session; fee-adjusted expectancy", Rule: "expectativa positiva DEPOIS de custos"},
		{Name: "demo_slippage_asia", Context: "Nikkei futures; Phase 1; volatile open; slippage-adjusted expectancy", Rule: "expectativa positiva DEPOIS de custos"},
		{Name: "demo_roll_cost_commodities", Context: "Commodity futures; Phase 1; contract roll; total trading costs", Rule: "expectativa positiva DEPOIS de custos"},
		{Name: "demo_weekend_crypto_cost", Context: "Crypto spot; Phase 1; weekend; spread and fee costs", Rule: "expectativa positiva DEPOIS de custos"},
		{Name: "positive_expectancy_equities", Context: "Global equity basket; Phase 1; quiet regime; net expectancy threshold", Rule: "expectativa positiva DEPOIS de custos"},
		{Name: "positive_expectancy_rates", Context: "Sovereign bond futures; Phase 1; rate shock; net expectancy threshold", Rule: "expectativa positiva DEPOIS de custos"},
		{Name: "positive_expectancy_energy", Context: "Energy futures; Phase 1; inventory release; net expectancy threshold", Rule: "expectativa positiva DEPOIS de custos"},
		{Name: "positive_expectancy_agri", Context: "Agricultural futures; Phase 1; weather gap; net expectancy threshold", Rule: "expectativa positiva DEPOIS de custos"},
		{Name: "positive_expectancy_fx", Context: "Emerging FX; Phase 1; intervention rumor; net expectancy threshold", Rule: "expectativa positiva DEPOIS de custos"},
		{Name: "shadow_veto_tick_us", Context: "US index; shadow ledger; cash session; veto tracked tick by tick", Rule: "Todo veto é liquidado tick a tick"},
		{Name: "shadow_veto_tick_eu", Context: "European shares; shadow ledger; close auction; veto tracked tick by tick", Rule: "Todo veto é liquidado tick a tick"},
		{Name: "shadow_veto_tick_asia", Context: "Asian futures; shadow ledger; overnight; veto tracked tick by tick", Rule: "Todo veto é liquidado tick a tick"},
		{Name: "shadow_veto_tick_commodity", Context: "Copper futures; shadow ledger; supply shock; veto tracked tick by tick", Rule: "Todo veto é liquidado tick a tick"},
		{Name: "shadow_veto_tick_crypto", Context: "Crypto pair; shadow ledger; weekend; veto tracked tick by tick", Rule: "Todo veto é liquidado tick a tick"},
		{Name: "shadow_real_spread_fx", Context: "EUR/USD; shadow ledger; London open; actual spread recorded", Rule: "Todo veto é liquidado tick a tick"},
		{Name: "shadow_real_spread_gap", Context: "US equities; shadow ledger; gap open; actual spread recorded", Rule: "Todo veto é liquidado tick a tick"},
		{Name: "no_closed_candle_us", Context: "US index futures; shadow ledger; fast market; avoid closed-candle proxy", Rule: "nunca OHLC de candle fechado."},
		{Name: "no_closed_candle_asia", Context: "Asia FX; shadow ledger; overnight; avoid closed-candle proxy", Rule: "nunca OHLC de candle fechado."},
		{Name: "no_lookahead_election", Context: "Sovereign FX; election night; prospective evidence only", Rule: "nunca OHLC de candle fechado."},
		{Name: "score_losses_avoided", Context: "Global equities; weekly ledger; avoided-loss component measured", Rule: "Duas semanas seguidas com escore simétrico acima da barra da seção 3"},
		{Name: "score_missed_gains", Context: "Global FX; weekly ledger; missed-gain component measured", Rule: "Duas semanas seguidas com escore simétrico acima da barra da seção 3"},
		{Name: "rr_asymmetry_one_to_two", Context: "Index futures; ledger review; 1:2 risk/reward threshold context", Rule: "Duas semanas seguidas com escore simétrico acima da barra da seção 3"},
		{Name: "two_week_ai_retention_us", Context: "US equities; Phase 2; week two; score above bar", Rule: "Duas semanas seguidas com escore simétrico acima da barra da seção 3"},
		{Name: "two_week_ai_retention_fx", Context: "G10 FX; Phase 2; week two; score above bar", Rule: "Duas semanas seguidas com escore simétrico acima da barra da seção 3"},
		{Name: "two_week_ai_retention_rates", Context: "Rates futures; Phase 2; week two; score above bar", Rule: "Duas semanas seguidas com escore simétrico acima da barra da seção 3"},
		{Name: "ai_disable_under_bar_commodities", Context: "Energy futures; Phase 2; two weeks below bar; disable layer", Rule: "Duas semanas consecutivas abaixo da barra"},
		{Name: "ai_disable_under_bar_equities", Context: "European equities; Phase 2; two weeks below bar; disable layer", Rule: "Duas semanas consecutivas abaixo da barra"},
		{Name: "ai_disable_under_bar_crypto", Context: "Crypto; Phase 2; two weeks below bar; disable layer", Rule: "Duas semanas consecutivas abaixo da barra"},
		{Name: "second_agent_evidence_only", Context: "Global macro basket; Phase 3; ledger evidence required", Rule: "Somente com evidência do ledger da Fase 2 pedindo isso"},
		{Name: "new_source_evidence_gate", Context: "Cross-asset news; Phase 3; new source requires ledger evidence", Rule: "Somente com evidência do ledger da Fase 2 pedindo isso"},
		{Name: "tuning_evidence_gate", Context: "Global futures; Phase 3; tuning requires ledger evidence", Rule: "Somente com evidência do ledger da Fase 2 pedindo isso"},
		{Name: "phase_order_no_early_ai", Context: "US payroll event; before Phase 1 gate; AI remains out of scope", Rule: "≥ 200 sinais gravados"},
		{Name: "phase_order_shadow_first", Context: "Gold futures; Phase 1; no AI; shadow-only evidence", Rule: "≥ 200 sinais gravados"},
		{Name: "phase_order_crypto_weekend", Context: "Digital assets; Phase 1; weekend; demo-only gate", Rule: "cada limite em sandbox"},
	}
	contract.RunScenarios(t, document, scenarios)
}
