package development_test

import (
	"testing"

	"botfinac/tests/internal/contract"
)

func TestDevelopmentGlobalMarketScenarios(t *testing.T) {
	document, err := contract.ReadArchitecture()
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []contract.Scenario{
		{Name: "fed_rate_decision", Context: "US equities; North America; cash session; rate surprise; official release", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "us_cpi_gap", Context: "US equity index futures; North America; pre-open; inflation gap; official statistics", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "us_payrolls_volatility", Context: "USD FX; North America; release window; employment shock; official calendar", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "treasury_auction_tail", Context: "US sovereign futures; North America; auction window; yield spike; official data", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "sec_filing_earnings", Context: "US equities; North America; regular session; earnings gap; official filing", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "exchange_halt_headline", Context: "US equity index; North America; cash session; halt headline; source classification", Rule: "LLM traduz texto em enums, não decide."},
		{Name: "ecb_policy_surprise", Context: "EUR FX; Europe; policy window; rate surprise; official release", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "euro_area_inflation", Context: "European equities; Europe; pre-open; CPI gap; official data", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "boe_emergency_statement", Context: "GBP FX; Europe; overnight; emergency policy; sanitized notice", Rule: "Lê texto sanitizado dentro de tags <untrusted>"},
		{Name: "sovereign_spread_widening", Context: "Euro sovereign bonds; Europe; regular session; spread shock; official bulletin", Rule: "Sanitização → fila com prioridade (oficial > ruído) → agente LLM."},
		{Name: "snb_fx_comment", Context: "CHF FX; Europe; London session; intervention rumor; source classification", Rule: "Fonte de ruído → HOLD só com anomalia de volume no WebSocket."},
		{Name: "swiss_cpi_revision", Context: "CHF rates; Europe; release window; data revision; official source", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "boj_yield_policy", Context: "JPY FX; Asia-Pacific; Tokyo session; policy shift; official release", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "tokyo_equity_open_gap", Context: "Japanese equities; Asia-Pacific; cash open; gap; event context", Rule: "LLM traduz texto em enums, não decide."},
		{Name: "pboc_fix_surprise", Context: "CNH FX; Asia-Pacific; Asia session; fixing surprise; official source", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "china_pmi_shock", Context: "Asian equity index futures; Asia-Pacific; pre-open; PMI shock; official data", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "hong_kong_listing_notice", Context: "Hong Kong equities; Asia-Pacific; cash session; issuer notice; sanitized feed", Rule: "Lê texto sanitizado dentro de tags <untrusted>"},
		{Name: "rba_rate_decision", Context: "AUD FX; Asia-Pacific; policy window; rate surprise; official release", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "india_inflation_release", Context: "INR FX; Asia-Pacific; release window; CPI surprise; official data", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "india_market_rumor", Context: "Indian equities; Asia-Pacific; regular session; unverified rumor; volume anomaly check", Rule: "Fonte de ruído → HOLD só com anomalia de volume no WebSocket."},
		{Name: "copom_rate_surprise", Context: "BRL FX; South America; policy window; rate surprise; official notice", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "cvm_material_fact", Context: "Brazilian equities; South America; regular session; material fact; official filing", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "bcb_emergency_notice", Context: "BRL rates; South America; overnight; emergency bulletin; sanitized text", Rule: "Lê texto sanitizado dentro de tags <untrusted>"},
		{Name: "brl_social_rumor", Context: "BRL FX; South America; London-New York overlap; social rumor; noise source", Rule: "Fonte de ruído → HOLD só com anomalia de volume no WebSocket."},
		{Name: "brazil_earnings_gap", Context: "Brazilian equities; South America; cash open; earnings gap; official filing", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "banxico_policy_statement", Context: "MXN FX; North America; policy window; rate surprise; official release", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "bank_of_canada_outlook", Context: "CAD FX; North America; regular session; guidance shock; official statement", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "latam_bond_auction", Context: "Emerging-market sovereign debt; Americas; auction window; auction result; official data", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "opec_production_notice", Context: "Crude oil futures; Middle East/global; overnight; supply shock; official notice", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "eia_inventory_surprise", Context: "Energy futures; North America; release window; inventory surprise; official data", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "lng_supply_disruption", Context: "Natural gas; Europe/Asia; overnight; infrastructure event; sanitized source", Rule: "Lê texto sanitizado dentro de tags <untrusted>"},
		{Name: "gold_geopolitical_gap", Context: "Gold futures; global; overnight; geopolitical shock; event classification", Rule: "LLM traduz texto em enums, não decide."},
		{Name: "copper_mine_disruption", Context: "Copper futures; South America/global; regular session; supply event; source priority", Rule: "Sanitização → fila com prioridade (oficial > ruído) → agente LLM."},
		{Name: "agriculture_weather_alert", Context: "Agricultural futures; Americas; overnight; weather alert; official bulletin", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "freight_route_blockage", Context: "Global commodity baskets; cross-region; overnight; route disruption; sanitized feed", Rule: "Lê texto sanitizado dentro de tags <untrusted>"},
		{Name: "crypto_etf_filing", Context: "Crypto assets; North America/global; continuous; regulator filing; official source", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "crypto_exchange_outage", Context: "Crypto assets; global; weekend; venue outage; fail-closed classification", Rule: "budget esgotado/timeout/análise inválida → degrada para SHOCK_WAIT"},
		{Name: "stablecoin_depeg_report", Context: "Digital assets; global; weekend; depeg report; untrusted-text boundary", Rule: "Lê texto sanitizado dentro de tags <untrusted>"},
		{Name: "crypto_weekend_liquidity", Context: "Digital assets; global; weekend; thin liquidity; deterministic path", Rule: "Todo sinal nasce do motor determinístico. LLM nunca toca o caminho crítico da ordem."},
		{Name: "futures_roll_calendar", Context: "Global index futures; cross-region; roll window; calendar context; enum classification", Rule: "LLM traduz texto em enums, não decide."},
		{Name: "cross_asset_correlation_break", Context: "FX and equity indices; global; overlapping sessions; correlation shock; deterministic math", Rule: "Todo sinal nasce do motor determinístico. LLM nunca toca o caminho crítico da ordem."},
		{Name: "flash_crash_official_alert", Context: "Multi-asset indices; global; regular session; flash crash; official alert", Rule: "Fonte oficial e categoria relevante → HOLD em 1 ms."},
		{Name: "central_bank_press_conference", Context: "Global FX; cross-region; live session; policy language; enum-only output", Rule: "LLM traduz texto em enums, não decide."},
		{Name: "election_result_gap", Context: "Sovereign rates and FX; global; overnight; election shock; sanitized sources", Rule: "Lê texto sanitizado dentro de tags <untrusted>"},
		{Name: "cyberattack_exchange_notice", Context: "Multi-market instruments; global; overnight; cyber outage notice; official priority", Rule: "Sanitização → fila com prioridade (oficial > ruído) → agente LLM."},
		{Name: "duplicate_rss_headlines", Context: "Cross-asset watchlist; global; overlapping sessions; duplicated news; dedup pipeline", Rule: "Sanitização → fila com prioridade (oficial > ruído) → agente LLM."},
		{Name: "prompt_injection_in_headline", Context: "Cross-asset news; global; any session; hostile text; untrusted boundary", Rule: "Lê texto sanitizado dentro de tags <untrusted>"},
		{Name: "invalid_llm_schema", Context: "Multi-asset watchlist; global; event window; invalid enum response; fail closed", Rule: "Agente responde em enums dentro do prazo (8–10 s); estourou → degrada para SHOCK_WAIT."},
		{Name: "llm_timeout_payroll", Context: "USD FX and rates; North America; payroll window; model timeout; degrade conservatively", Rule: "Agente responde em enums dentro do prazo (8–10 s); estourou → degrada para SHOCK_WAIT."},
		{Name: "llm_budget_exhausted_copom", Context: "BRL FX; South America; policy window; model budget exhausted; conservative state", Rule: "budget esgotado/timeout/análise inválida → degrada para SHOCK_WAIT"},
		{Name: "official_feed_silence", Context: "Global macro calendar; cross-region; scheduled release; silent ingestion; heartbeat risk", Rule: "Agente responde em enums dentro do prazo (8–10 s); estourou → degrada para SHOCK_WAIT."},
	}
	contract.RunScenarios(t, document, scenarios)
}
