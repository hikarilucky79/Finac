package production_test

import (
	"testing"

	"botfinac/tests/internal/contract"
)

func TestProductionGlobalMarketScenarios(t *testing.T) {
	document, err := contract.ReadArchitecture()
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []contract.Scenario{
		{Name: "equity_fill_disconnect", Context: "US equities; North America; cash open; fill during broker disconnect", Rule: "motor fica congelado até a posição local ser reconciliada com a da corretora."},
		{Name: "fx_partial_fill_reconnect", Context: "EUR/USD; Europe; London session; partial fill during reconnect", Rule: "motor fica congelado até a posição local ser reconciliada com a da corretora."},
		{Name: "futures_gap_reconcile", Context: "S&P futures; North America; overnight gap; reconcile broker position", Rule: "motor fica congelado até a posição local ser reconciliada com a da corretora."},
		{Name: "yen_flash_crash_reconcile", Context: "JPY FX; Asia-Pacific; Tokyo session; flash crash and reconnect", Rule: "motor fica congelado até a posição local ser reconciliada com a da corretora."},
		{Name: "crypto_weekend_restart", Context: "Crypto perpetuals; global; weekend; process restart with open position", Rule: "Crash/restart de qualquer processo"},
		{Name: "oil_inventory_feed_loss", Context: "Crude futures; North America; inventory release; market feed lost", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "gold_geo_shock_llm_timeout", Context: "Gold futures; global; overnight; model timeout during shock", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "euro_auction_close_stop", Context: "European equities; Europe; closing auction; stop modification", Rule: "amend para apertar stop (nunca cancel+new)"},
		{Name: "nasdaq_volatility_oco", Context: "Nasdaq futures; North America; cash session; protective OCO remains broker-side", Rule: "OCO residente na corretora"},
		{Name: "brl_copom_duplicate_resend", Context: "BRL FX; South America; policy window; duplicate order resend", Rule: "ordem OCO envolvente dispara para a corretora com chave de idempotência."},
		{Name: "crypto_venue_duplicate_fill", Context: "Crypto spot; global; weekend; duplicate fill notification", Rule: "ordem OCO envolvente dispara para a corretora com chave de idempotência."},
		{Name: "copper_opening_auction", Context: "Copper futures; global; exchange auction; wait for liquidity window", Rule: "espera leilão"},
		{Name: "agri_limit_move_protection", Context: "Agricultural futures; Americas; limit-move context; no naked market exit", Rule: "nunca market nu"},
		{Name: "eur_bond_state_store_down", Context: "Euro sovereign futures; Europe; regular session; state store unavailable", Rule: "Store de estados inacessível"},
		{Name: "em_fx_heartbeat_silent", Context: "Emerging FX; global; overnight; ingestion heartbeat silent", Rule: "Ingestion sem heartbeat > X min"},
		{Name: "asia_equity_heartbeat_gap", Context: "Asia equities; Asia-Pacific; open; news ingestion heartbeat lost", Rule: "Ingestion sem heartbeat > X min"},
		{Name: "us_holiday_feed_silence", Context: "US rates; North America; holiday session; sensor silent", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "boe_veto_restart", Context: "GBP rates; Europe; policy announcement; persisted directional veto on restart", Rule: "Crash/restart de qualquer processo"},
		{Name: "boj_veto_ttl_clock", Context: "JPY FX; Asia-Pacific; Tokyo session; veto TTL uses clock timestamp", Rule: "Crash/restart de qualquer processo"},
		{Name: "ecb_schema_invalid", Context: "EUR FX; Europe; ECB window; invalid model schema", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "fed_budget_exhaustion", Context: "US rates; North America; Fed window; model budget exhausted", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "opec_analysis_timeout", Context: "Energy futures; global; overnight; classifier timeout", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "cvm_feed_sanitization_fault", Context: "Brazilian equities; South America; cash session; untrusted raw headline", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "rba_source_noise_volume", Context: "AUD FX; Asia-Pacific; release window; noisy source without volume anomaly", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "sec_official_shock_hold", Context: "US equities; North America; regular session; official shock requires hold", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "election_gap_human_kill", Context: "Sovereign FX; global; election night; operator activates human kill switch", Rule: "expõe o kill switch manual"},
		{Name: "cyberattack_guardian_survival", Context: "Multi-asset; global; overnight; main stack failure; guardian remains independent", Rule: "Processo separado, independente"},
		{Name: "daily_equity_stop_rates", Context: "Global rates; cross-region; end of day; daily equity stop disconnects", Rule: "equity stop (perda total do dia = zero e desconecta)"},
		{Name: "crypto_margin_loss_limit", Context: "Crypto derivatives; global; weekend; daily loss reaches zero threshold", Rule: "equity stop (perda total do dia = zero e desconecta)"},
		{Name: "index_futures_max_trades", Context: "Equity index futures; North America; volatile session; max trades hard rule", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "fx_minimum_rr_veto", Context: "G10 FX; London-New York overlap; low reward-risk signal veto", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "commodity_cooldown_after_loss", Context: "Commodity futures; global; post-loss session; cooldown remains active", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "market_hours_boundary", Context: "US index; North America; session close; time gate enforced", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "state_veto_blocks_new_entry", Context: "Asian equity futures; Asia-Pacific; event shock; state machine veto", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "position_reconcile_before_order", Context: "Multi-asset; global; reconnect; no order before broker reconciliation", Rule: "motor fica congelado até a posição local ser reconciliada com a da corretora."},
		{Name: "open_oco_survives_process_crash", Context: "Crude futures; global; overnight; broker-resident OCO after crash", Rule: "OCO residente na corretora"},
		{Name: "stop_amend_no_cancel_replace", Context: "Gold futures; global; fast market; amend stop without cancel-new", Rule: "amend para apertar stop (nunca cancel+new)"},
		{Name: "protective_exit_illiquid_market", Context: "Front-month futures; cross-region; liquidity vacuum; bounded protective exit", Rule: "nunca market nu"},
		{Name: "auction_exit_wait", Context: "European equities; Europe; closing auction; wait for auction state", Rule: "espera leilão"},
		{Name: "directional_veto_per_side", Context: "BRL FX; South America; policy shock; block affected side only", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "hold_ttl_expiry_explicit", Context: "US equities; North America; post-news; return to normal only by explicit TTL", Rule: "HOLD **jamais** pode expirar para NORMAL"},
		{Name: "shock_wait_not_normal", Context: "Global FX; cross-session; event still unresolved; hold cannot expire to normal", Rule: "HOLD **jamais** pode expirar para NORMAL"},
		{Name: "stale_position_cache", Context: "US index options context; North America; reconnect; broker position is source of truth", Rule: "Posição local é cache, nunca verdade."},
		{Name: "guardian_independent_config", Context: "Global assets; all sessions; corrupted main config; separate guardian config", Rule: "Processo separado, independente"},
		{Name: "guardian_reads_broker_direct", Context: "FX basket; global; broker reconnect; guardian reads broker balances directly", Rule: "Processo separado, independente"},
		{Name: "ledger_failure_nonblocking", Context: "Cross-asset; global; ledger unavailable; ledger not operational dependency", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
		{Name: "crash_read_persisted_state", Context: "Commodity basket; global; restart after state transition; reload persistence", Rule: "Crash/restart de qualquer processo"},
		{Name: "idempotent_retry_after_timeout", Context: "Nasdaq futures; North America; order timeout; retry uses idempotency key", Rule: "ordem OCO envolvente dispara para a corretora com chave de idempotência."},
		{Name: "liquidation_after_veto_liquidity", Context: "European index futures; Europe; thin liquidity; exit at next liquid window", Rule: "espera leilão"},
		{Name: "websocket_drop_during_fill", Context: "FX futures; Asia-Pacific; disconnect while fill is executing", Rule: "motor fica congelado até a posição local ser reconciliada com a da corretora."},
		{Name: "market_wide_circuit_breaker", Context: "Global equity index; regional session; volatility interruption; fail closed", Rule: "Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho."},
	}
	contract.RunScenarios(t, document, scenarios)
}
