package accessor

import "time"

type GetAllAssetsResponse struct {
	Status struct {
		Timestamp time.Time `json:"timestamp"`
		Elapsed   int       `json:"elapsed"`
	} `json:"status"`
	Data []struct {
		ID      string `json:"id"`
		Symbol  string `json:"symbol"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		Profile struct {
			Profile struct {
				General struct {
					Overview struct {
						IsVerified     bool   `json:"is_verified"`
						Tagline        string `json:"tagline"`
						Category       string `json:"category"`
						Sector         string `json:"sector"`
						Tags           string `json:"tags"`
						ProjectDetails string `json:"project_details"`
						OfficialLinks  []struct {
							Name string `json:"name"`
							Link string `json:"link"`
						} `json:"official_links"`
					} `json:"overview"`
					Background struct {
						BackgroundDetails    string `json:"background_details"`
						IssuingOrganizations []struct {
							Slug        string `json:"slug"`
							Name        string `json:"name"`
							Logo        string `json:"logo"`
							Description string `json:"description"`
						} `json:"issuing_organizations"`
					} `json:"background"`
					Roadmap []struct {
						Title   string    `json:"title"`
						Date    time.Time `json:"date"`
						Type    string    `json:"type"`
						Details string    `json:"details"`
					} `json:"roadmap"`
					Regulation struct {
						RegulatoryDetails string `json:"regulatory_details"`
						SfarScore         int    `json:"sfar_score"`
						SfarSummary       string `json:"sfar_summary"`
					} `json:"regulation"`
				} `json:"general"`
				Contributors struct {
					Individuals []struct {
						Slug        string `json:"slug"`
						FirstName   string `json:"first_name"`
						LastName    string `json:"last_name"`
						Title       string `json:"title"`
						Description string `json:"description"`
						AvatarURL   string `json:"avatar_url"`
					} `json:"individuals"`
					Organizations []struct {
						Slug        string `json:"slug"`
						Name        string `json:"name"`
						Logo        string `json:"logo"`
						Description string `json:"description"`
					} `json:"organizations"`
				} `json:"contributors"`
				Advisors struct {
					Individuals []struct {
						Slug        string `json:"slug"`
						FirstName   string `json:"first_name"`
						LastName    string `json:"last_name"`
						Title       string `json:"title"`
						Description string `json:"description"`
						AvatarURL   string `json:"avatar_url"`
					} `json:"individuals"`
					Organizations []struct {
						Slug        string `json:"slug"`
						Name        string `json:"name"`
						Logo        string `json:"logo"`
						Description string `json:"description"`
					} `json:"organizations"`
				} `json:"advisors"`
				Investors struct {
					Individuals []struct {
						Slug        string `json:"slug"`
						FirstName   string `json:"first_name"`
						LastName    string `json:"last_name"`
						Title       string `json:"title"`
						Description string `json:"description"`
						AvatarURL   string `json:"avatar_url"`
					} `json:"individuals"`
					Organizations []struct {
						Slug        string `json:"slug"`
						Name        string `json:"name"`
						Logo        string `json:"logo"`
						Description string `json:"description"`
					} `json:"organizations"`
				} `json:"investors"`
				Ecosystem struct {
					Assets []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
					} `json:"assets"`
					Organizations []struct {
						Slug        string `json:"slug"`
						Name        string `json:"name"`
						Logo        string `json:"logo"`
						Description string `json:"description"`
					} `json:"organizations"`
				} `json:"ecosystem"`
				Economics struct {
					Token struct {
						TokenName      string `json:"token_name"`
						TokenType      string `json:"token_type"`
						BlockExplorers []struct {
							Name string `json:"name"`
							Link string `json:"link"`
						} `json:"block_explorers"`
						Multitoken []struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"multitoken"`
						TokenUsage                  string `json:"token_usage"`
						TokenUsageDetailsAndWallets string `json:"token_usage_details_and_wallets"`
					} `json:"token"`
					Launch struct {
						General struct {
							LaunchStyle   string `json:"launch_style"`
							LaunchDetails string `json:"launch_details"`
						} `json:"general"`
						Fundraising struct {
							SalesRounds []struct {
								RestrictedJurisdictions []interface{} `json:"restricted_jurisdictions"`
							} `json:"sales_rounds"`
							SalesDocuments []struct {
							} `json:"sales_documents"`
							SalesTreasuryAccounts []struct {
							} `json:"sales_treasury_accounts"`
							TreasuryPolicies []struct {
							} `json:"treasury_policies"`
							ProjectedUseOfSalesProceeds []struct {
							} `json:"projected_use_of_sales_proceeds"`
						} `json:"fundraising"`
						InitialDistribution struct {
							InitialSupply            float64 `json:"initial_supply"`
							InitialSupplyRepartition struct {
								AllocatedToInvestorsPercentage                 float64 `json:"allocated_to_investors_percentage"`
								AllocatedToOrganizationOrFoundersPercentage    float64 `json:"allocated_to_organization_or_founders_percentage"`
								AllocatedToPreminedRewardsOrAirdropsPercentage int     `json:"allocated_to_premined_rewards_or_airdrops_percentage"`
							} `json:"initial_supply_repartition"`
							TokenDistributionDate time.Time `json:"token_distribution_date"`
							GenesisBlockDate      time.Time `json:"genesis_block_date"`
						} `json:"initial_distribution"`
					} `json:"launch"`
					ConsensusAndEmission struct {
						Supply struct {
							SupplyCurveDetails  string `json:"supply_curve_details"`
							GeneralEmissionType string `json:"general_emission_type"`
							PreciseEmissionType string `json:"precise_emission_type"`
							IsCappedSupply      bool   `json:"is_capped_supply"`
							MaxSupply           int    `json:"max_supply"`
						} `json:"supply"`
						Consensus struct {
							ConsensusDetails          string    `json:"consensus_details"`
							GeneralConsensusMechanism string    `json:"general_consensus_mechanism"`
							PreciseConsensusMechanism string    `json:"precise_consensus_mechanism"`
							TargetedBlockTime         int       `json:"targeted_block_time"`
							BlockReward               int       `json:"block_reward"`
							MiningAlgorithm           string    `json:"mining_algorithm"`
							NextHalvingDate           time.Time `json:"next_halving_date"`
							IsVictimOf51PercentAttack bool      `json:"is_victim_of_51_percent_attack"`
						} `json:"consensus"`
					} `json:"consensus_and_emission"`
					NativeTreasury struct {
						Accounts []struct {
							AccountType string `json:"account_type"`
							Addresses   []struct {
							} `json:"addresses"`
						} `json:"accounts"`
						TreasuryUsageDetails string `json:"treasury_usage_details"`
					} `json:"native_treasury"`
				} `json:"economics"`
				Technology struct {
					Overview struct {
						TechnologyDetails  string `json:"technology_details"`
						ClientRepositories []struct {
							Name        string `json:"name"`
							Link        string `json:"link"`
							LicenseType string `json:"license_type"`
						} `json:"client_repositories"`
					} `json:"overview"`
					Security struct {
						Audits []struct {
							Title   string    `json:"title"`
							Date    time.Time `json:"date"`
							Type    string    `json:"type"`
							Details string    `json:"details"`
						} `json:"audits"`
						KnownExploitsAndVulnerabilities []struct {
							Title   string    `json:"title"`
							Date    time.Time `json:"date"`
							Type    string    `json:"type"`
							Details string    `json:"details"`
						} `json:"known_exploits_and_vulnerabilities"`
					} `json:"security"`
				} `json:"technology"`
				Governance struct {
					GovernanceDetails string `json:"governance_details"`
					OnchainGovernance struct {
						OnchainGovernanceType    string `json:"onchain_governance_type"`
						OnchainGovernanceDetails string `json:"onchain_governance_details"`
						IsTreasuryDecentralized  bool   `json:"is_treasury_decentralized"`
					} `json:"onchain_governance"`
					Grants []struct {
						FundingOrganizations []struct {
						} `json:"funding_organizations"`
						GrantProgramDetails string `json:"grant_program_details"`
					} `json:"grants"`
				} `json:"governance"`
			} `json:"profile"`
		} `json:"profile"`
		Metrics struct {
			ID         string `json:"id"`
			Symbol     string `json:"symbol"`
			Name       string `json:"name"`
			Slug       string `json:"slug"`
			MarketData struct {
				PriceUsd                               float64 `json:"price_usd"`
				PriceBtc                               int     `json:"price_btc"`
				VolumeLast24Hours                      float64 `json:"volume_last_24_hours"`
				RealVolumeLast24Hours                  float64 `json:"real_volume_last_24_hours"`
				VolumeLast24HoursOverstatementMultiple float64 `json:"volume_last_24_hours_overstatement_multiple"`
				PercentChangeUsdLast24Hours            float64 `json:"percent_change_usd_last_24_hours"`
				PercentChangeBtcLast24Hours            int     `json:"percent_change_btc_last_24_hours"`
				OhlcvLast1Hour                         struct {
					Open   float64 `json:"open"`
					High   float64 `json:"high"`
					Low    float64 `json:"low"`
					Close  float64 `json:"close"`
					Volume float64 `json:"volume"`
				} `json:"ohlcv_last_1_hour"`
				OhlcvLast24Hour struct {
					Open   float64 `json:"open"`
					High   float64 `json:"high"`
					Low    float64 `json:"low"`
					Close  float64 `json:"close"`
					Volume float64 `json:"volume"`
				} `json:"ohlcv_last_24_hour"`
			} `json:"market_data"`
			Supply struct {
				Y2050                float64 `json:"y_2050"`
				Y2050PercentIssued   int     `json:"y_2050_percent_issued"`
				SupplyYplus10        float64 `json:"supply_yplus_10"`
				YPlus10IssuedPercent int     `json:"y_plus10_issued_percent"`
				Liquid               int     `json:"liquid"`
				Circulating          int     `json:"circulating"`
				StockToFlow          int     `json:"stock_to_flow"`
			} `json:"supply"`
			BlockchainStats24Hours struct {
				TransactionVolume      float64     `json:"transaction_volume"`
				Nvt                    float64     `json:"nvt"`
				SumOfFees              float64     `json:"sum_of_fees"`
				MedianTxValue          float64     `json:"median_tx_value"`
				MedianTxFee            float64     `json:"median_tx_fee"`
				CountOfActiveAddresses int         `json:"count_of_active_addresses"`
				CountOfTx              int         `json:"count_of_tx"`
				CountOfPayments        int         `json:"count_of_payments"`
				NewIssuance            float64     `json:"new_issuance"`
				AverageDifficulty      int64       `json:"average_difficulty"`
				KilobytesAdded         float64     `json:"kilobytes_added"`
				CountOfBlocksAdded     int         `json:"count_of_blocks_added"`
				SupplyMovedOffChain    interface{} `json:"supply_moved_off_chain"`
			} `json:"blockchain_stats_24_hours"`
			AllTimeHigh struct {
				Price       int       `json:"price"`
				At          time.Time `json:"at"`
				DaysSince   int       `json:"days_since"`
				PercentDown float64   `json:"percent_down"`
			} `json:"all_time_high"`
			DeveloperActivity struct {
				Stars                   int `json:"stars"`
				Watchers                int `json:"watchers"`
				CommitsLast3Months      int `json:"commits_last_3_months"`
				CommitsLast1Year        int `json:"commits_last_1_year"`
				LinesAddedLast3Months   int `json:"lines_added_last_3_months"`
				LinesAddedLast1Year     int `json:"lines_added_last_1_year"`
				LinesDeletedLast3Months int `json:"lines_deleted_last_3_months"`
				LinesDeletedLast1Year   int `json:"lines_deleted_last_1_year"`
			} `json:"developer_activity"`
			RoiData struct {
				PercentChangeLast1Week   float64 `json:"percent_change_last_1_week"`
				PercentChangeLast1Month  float64 `json:"percent_change_last_1_month"`
				PercentChangeLast3Months float64 `json:"percent_change_last_3_months"`
				PercentChangeLast1Year   float64 `json:"percent_change_last_1_year"`
			} `json:"roi_data"`
			MiscData struct {
				AssetAgeDays     int        `json:"asset_age_days"`
				VladimirClubCost float64    `json:"vladimir_club_cost"`
				Categories       [][]string `json:"categories"`
				Sector           [][]string `json:"sector"`
			} `json:"misc_data"`
		} `json:"metrics"`
	} `json:"data"`
}

type GetAssetResponse struct {
	Status struct {
		Timestamp time.Time `json:"timestamp"`
		Elapsed   int       `json:"elapsed"`
	} `json:"status"`
	Data struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
		Slug   string `json:"slug"`
	} `json:"data"`
}
type GetProfileResponse struct {
	Status struct {
		Timestamp time.Time `json:"timestamp"`
		Elapsed   int       `json:"elapsed"`
	} `json:"status"`
	Data struct {
		Profile struct {
			General struct {
				Overview struct {
					IsVerified     bool   `json:"is_verified"`
					Tagline        string `json:"tagline"`
					Category       string `json:"category"`
					Sector         string `json:"sector"`
					Tags           string `json:"tags"`
					ProjectDetails string `json:"project_details"`
					OfficialLinks  []struct {
						Name string `json:"name"`
						Link string `json:"link"`
					} `json:"official_links"`
				} `json:"overview"`
				Background struct {
					BackgroundDetails    string `json:"background_details"`
					IssuingOrganizations []struct {
						Slug        string `json:"slug"`
						Name        string `json:"name"`
						Logo        string `json:"logo"`
						Description string `json:"description"`
					} `json:"issuing_organizations"`
				} `json:"background"`
				Roadmap []struct {
					Title   string    `json:"title"`
					Date    time.Time `json:"date"`
					Type    string    `json:"type"`
					Details string    `json:"details"`
				} `json:"roadmap"`
				Regulation struct {
					RegulatoryDetails string `json:"regulatory_details"`
					SfarScore         int    `json:"sfar_score"`
					SfarSummary       string `json:"sfar_summary"`
				} `json:"regulation"`
			} `json:"general"`
			Contributors struct {
				Individuals []struct {
					Slug        string `json:"slug"`
					FirstName   string `json:"first_name"`
					LastName    string `json:"last_name"`
					Title       string `json:"title"`
					Description string `json:"description"`
					AvatarURL   string `json:"avatar_url"`
				} `json:"individuals"`
				Organizations []struct {
					Slug        string `json:"slug"`
					Name        string `json:"name"`
					Logo        string `json:"logo"`
					Description string `json:"description"`
				} `json:"organizations"`
			} `json:"contributors"`
			Advisors struct {
				Individuals []struct {
					Slug        string `json:"slug"`
					FirstName   string `json:"first_name"`
					LastName    string `json:"last_name"`
					Title       string `json:"title"`
					Description string `json:"description"`
					AvatarURL   string `json:"avatar_url"`
				} `json:"individuals"`
				Organizations []struct {
					Slug        string `json:"slug"`
					Name        string `json:"name"`
					Logo        string `json:"logo"`
					Description string `json:"description"`
				} `json:"organizations"`
			} `json:"advisors"`
			Investors struct {
				Individuals []struct {
					Slug        string `json:"slug"`
					FirstName   string `json:"first_name"`
					LastName    string `json:"last_name"`
					Title       string `json:"title"`
					Description string `json:"description"`
					AvatarURL   string `json:"avatar_url"`
				} `json:"individuals"`
				Organizations []struct {
					Slug        string `json:"slug"`
					Name        string `json:"name"`
					Logo        string `json:"logo"`
					Description string `json:"description"`
				} `json:"organizations"`
			} `json:"investors"`
			Ecosystem struct {
				Assets []struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"assets"`
				Organizations []struct {
					Slug        string `json:"slug"`
					Name        string `json:"name"`
					Logo        string `json:"logo"`
					Description string `json:"description"`
				} `json:"organizations"`
			} `json:"ecosystem"`
			Economics struct {
				Token struct {
					TokenName      string `json:"token_name"`
					TokenType      string `json:"token_type"`
					BlockExplorers []struct {
						Name string `json:"name"`
						Link string `json:"link"`
					} `json:"block_explorers"`
					Multitoken []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
					} `json:"multitoken"`
					TokenUsage                  string `json:"token_usage"`
					TokenUsageDetailsAndWallets string `json:"token_usage_details_and_wallets"`
				} `json:"token"`
				Launch struct {
					General struct {
						LaunchStyle   string `json:"launch_style"`
						LaunchDetails string `json:"launch_details"`
					} `json:"general"`
					Fundraising struct {
						SalesRounds []struct {
							Title                        string    `json:"title"`
							StartDate                    time.Time `json:"start_date"`
							Type                         string    `json:"type"`
							Details                      string    `json:"details"`
							EndDate                      time.Time `json:"end_date"`
							NativeTokensAllocated        int       `json:"native_tokens_allocated"`
							AssetCollected               string    `json:"asset_collected"`
							PricePerTokenInAsset         float64   `json:"price_per_token_in_asset"`
							EquivalentPricePerTokenInUSD float64   `json:"equivalent_price_per_token_in_USD"`
							AmountCollectedInAsset       int       `json:"amount_collected_in_asset"`
							AmountCollectedInUSD         int       `json:"amount_collected_in_USD"`
							IsKycRequired                bool      `json:"is_kyc_required"`
							RestrictedJurisdictions      []string  `json:"restricted_jurisdictions"`
						} `json:"sales_rounds"`
						SalesDocuments []struct {
							Name string `json:"name"`
							Link string `json:"link"`
						} `json:"sales_documents"`
						SalesTreasuryAccounts []struct {
							AccountType string `json:"account_type"`
							AssetHeld   struct {
								ID   string `json:"id"`
								Name string `json:"name"`
							} `json:"asset_held"`
							Addresses []struct {
							} `json:"addresses"`
							Security string `json:"security"`
						} `json:"sales_treasury_accounts"`
						TreasuryPolicies []struct {
							Name string `json:"name"`
							Link string `json:"link"`
						} `json:"treasury_policies"`
						ProjectedUseOfSalesProceeds []struct {
							Category           string  `json:"category"`
							AmountInPercentage float64 `json:"amount_in_percentage"`
						} `json:"projected_use_of_sales_proceeds"`
					} `json:"fundraising"`
					InitialDistribution struct {
						InitialSupply            float64 `json:"initial_supply"`
						InitialSupplyRepartition struct {
							AllocatedToInvestorsPercentage                 float64 `json:"allocated_to_investors_percentage"`
							AllocatedToOrganizationOrFoundersPercentage    float64 `json:"allocated_to_organization_or_founders_percentage"`
							AllocatedToPreminedRewardsOrAirdropsPercentage int     `json:"allocated_to_premined_rewards_or_airdrops_percentage"`
						} `json:"initial_supply_repartition"`
						TokenDistributionDate time.Time `json:"token_distribution_date"`
						GenesisBlockDate      time.Time `json:"genesis_block_date"`
					} `json:"initial_distribution"`
				} `json:"launch"`
				ConsensusAndEmission struct {
					Supply struct {
						SupplyCurveDetails  string `json:"supply_curve_details"`
						GeneralEmissionType string `json:"general_emission_type"`
						PreciseEmissionType string `json:"precise_emission_type"`
						IsCappedSupply      bool   `json:"is_capped_supply"`
						MaxSupply           int    `json:"max_supply"`
					} `json:"supply"`
					Consensus struct {
						ConsensusDetails          string    `json:"consensus_details"`
						GeneralConsensusMechanism string    `json:"general_consensus_mechanism"`
						PreciseConsensusMechanism string    `json:"precise_consensus_mechanism"`
						TargetedBlockTime         int       `json:"targeted_block_time"`
						BlockReward               int       `json:"block_reward"`
						MiningAlgorithm           string    `json:"mining_algorithm"`
						NextHalvingDate           time.Time `json:"next_halving_date"`
						IsVictimOf51PercentAttack bool      `json:"is_victim_of_51_percent_attack"`
					} `json:"consensus"`
				} `json:"consensus_and_emission"`
				NativeTreasury struct {
					Accounts []struct {
						AccountType string `json:"account_type"`
						Addresses   []struct {
							Name string `json:"name"`
							Link string `json:"link"`
						} `json:"addresses"`
					} `json:"accounts"`
					TreasuryUsageDetails string `json:"treasury_usage_details"`
				} `json:"native_treasury"`
			} `json:"economics"`
			Technology struct {
				Overview struct {
					TechnologyDetails  string `json:"technology_details"`
					ClientRepositories []struct {
						Name        string `json:"name"`
						Link        string `json:"link"`
						LicenseType string `json:"license_type"`
					} `json:"client_repositories"`
				} `json:"overview"`
				Security struct {
					Audits []struct {
						Title   string    `json:"title"`
						Date    time.Time `json:"date"`
						Type    string    `json:"type"`
						Details string    `json:"details"`
					} `json:"audits"`
					KnownExploitsAndVulnerabilities []struct {
						Title   string    `json:"title"`
						Date    time.Time `json:"date"`
						Type    string    `json:"type"`
						Details string    `json:"details"`
					} `json:"known_exploits_and_vulnerabilities"`
				} `json:"security"`
			} `json:"technology"`
			Governance struct {
				GovernanceDetails string `json:"governance_details"`
				OnchainGovernance struct {
					OnchainGovernanceType    string `json:"onchain_governance_type"`
					OnchainGovernanceDetails string `json:"onchain_governance_details"`
					IsTreasuryDecentralized  bool   `json:"is_treasury_decentralized"`
				} `json:"onchain_governance"`
				Grants []struct {
					FundingOrganizations []struct {
						Slug        string `json:"slug"`
						Name        string `json:"name"`
						Logo        string `json:"logo"`
						Description string `json:"description"`
					} `json:"funding_organizations"`
					GrantProgramDetails string `json:"grant_program_details"`
				} `json:"grants"`
			} `json:"governance"`
		} `json:"profile"`
	} `json:"data"`
}

type GetMetricsResponse struct {
	Status struct {
		Timestamp time.Time `json:"timestamp"`
		Elapsed   int       `json:"elapsed"`
	} `json:"status"`
	Data struct {
		ID         string `json:"id"`
		Symbol     string `json:"symbol"`
		Name       string `json:"name"`
		Slug       string `json:"slug"`
		MarketData struct {
			PriceUsd                               float64 `json:"price_usd"`
			PriceBtc                               int     `json:"price_btc"`
			VolumeLast24Hours                      float64 `json:"volume_last_24_hours"`
			RealVolumeLast24Hours                  float64 `json:"real_volume_last_24_hours"`
			VolumeLast24HoursOverstatementMultiple float64 `json:"volume_last_24_hours_overstatement_multiple"`
			PercentChangeUsdLast24Hours            float64 `json:"percent_change_usd_last_24_hours"`
			PercentChangeBtcLast24Hours            int     `json:"percent_change_btc_last_24_hours"`
			OhlcvLast1Hour                         struct {
				Open   float64 `json:"open"`
				High   float64 `json:"high"`
				Low    float64 `json:"low"`
				Close  float64 `json:"close"`
				Volume float64 `json:"volume"`
			} `json:"ohlcv_last_1_hour"`
			OhlcvLast24Hour struct {
				Open   float64 `json:"open"`
				High   float64 `json:"high"`
				Low    float64 `json:"low"`
				Close  float64 `json:"close"`
				Volume float64 `json:"volume"`
			} `json:"ohlcv_last_24_hour"`
		} `json:"market_data"`
		Supply struct {
			Y2050                float64 `json:"y_2050"`
			Y2050PercentIssued   int     `json:"y_2050_percent_issued"`
			SupplyYplus10        float64 `json:"supply_yplus_10"`
			YPlus10IssuedPercent int     `json:"y_plus10_issued_percent"`
			Liquid               int     `json:"liquid"`
			Circulating          int     `json:"circulating"`
			StockToFlow          int     `json:"stock_to_flow"`
		} `json:"supply"`
		BlockchainStats24Hours struct {
			TransactionVolume      float64     `json:"transaction_volume"`
			Nvt                    float64     `json:"nvt"`
			SumOfFees              float64     `json:"sum_of_fees"`
			MedianTxValue          float64     `json:"median_tx_value"`
			MedianTxFee            float64     `json:"median_tx_fee"`
			CountOfActiveAddresses int         `json:"count_of_active_addresses"`
			CountOfTx              int         `json:"count_of_tx"`
			CountOfPayments        int         `json:"count_of_payments"`
			NewIssuance            float64     `json:"new_issuance"`
			AverageDifficulty      int64       `json:"average_difficulty"`
			KilobytesAdded         float64     `json:"kilobytes_added"`
			CountOfBlocksAdded     int         `json:"count_of_blocks_added"`
			SupplyMovedOffChain    interface{} `json:"supply_moved_off_chain"`
		} `json:"blockchain_stats_24_hours"`
		AllTimeHigh struct {
			Price       int       `json:"price"`
			At          time.Time `json:"at"`
			DaysSince   int       `json:"days_since"`
			PercentDown float64   `json:"percent_down"`
		} `json:"all_time_high"`
		DeveloperActivity struct {
			Stars                   int `json:"stars"`
			Watchers                int `json:"watchers"`
			CommitsLast3Months      int `json:"commits_last_3_months"`
			CommitsLast1Year        int `json:"commits_last_1_year"`
			LinesAddedLast3Months   int `json:"lines_added_last_3_months"`
			LinesAddedLast1Year     int `json:"lines_added_last_1_year"`
			LinesDeletedLast3Months int `json:"lines_deleted_last_3_months"`
			LinesDeletedLast1Year   int `json:"lines_deleted_last_1_year"`
		} `json:"developer_activity"`
		RoiData struct {
			PercentChangeLast1Week   float64 `json:"percent_change_last_1_week"`
			PercentChangeLast1Month  float64 `json:"percent_change_last_1_month"`
			PercentChangeLast3Months float64 `json:"percent_change_last_3_months"`
			PercentChangeLast1Year   float64 `json:"percent_change_last_1_year"`
		} `json:"roi_data"`
		MiscData struct {
			AssetAgeDays     int        `json:"asset_age_days"`
			VladimirClubCost float64    `json:"vladimir_club_cost"`
			Categories       [][]string `json:"categories"`
			Sector           [][]string `json:"sector"`
		} `json:"misc_data"`
	} `json:"data"`
}

type GetMarketDataResponse struct {
	Status struct {
		Timestamp time.Time `json:"timestamp"`
		Elapsed   int       `json:"elapsed"`
	} `json:"status"`
	Data struct {
		ID         string `json:"id"`
		Symbol     string `json:"symbol"`
		Name       string `json:"name"`
		Slug       string `json:"slug"`
		MarketData struct {
			PriceUsd                               float64 `json:"price_usd"`
			PriceBtc                               int     `json:"price_btc"`
			VolumeLast24Hours                      float64 `json:"volume_last_24_hours"`
			RealVolumeLast24Hours                  float64 `json:"real_volume_last_24_hours"`
			VolumeLast24HoursOverstatementMultiple float64 `json:"volume_last_24_hours_overstatement_multiple"`
			PercentChangeUsdLast24Hours            float64 `json:"percent_change_usd_last_24_hours"`
			PercentChangeBtcLast24Hours            int     `json:"percent_change_btc_last_24_hours"`
			OhlcvLast1Hour                         struct {
				Open   float64 `json:"open"`
				High   float64 `json:"high"`
				Low    float64 `json:"low"`
				Close  float64 `json:"close"`
				Volume float64 `json:"volume"`
			} `json:"ohlcv_last_1_hour"`
			OhlcvLast24Hour struct {
				Open   float64 `json:"open"`
				High   float64 `json:"high"`
				Low    float64 `json:"low"`
				Close  float64 `json:"close"`
				Volume float64 `json:"volume"`
			} `json:"ohlcv_last_24_hour"`
		} `json:"market_data"`
	} `json:"data"`
}

type GetTimeseriesResponse struct {
	Status struct {
		Elapsed   int       `json:"elapsed"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"status"`
	Data struct {
		ID                  string        `json:"id"`
		Symbol              string        `json:"symbol"`
		Name                string        `json:"name"`
		Slug                string        `json:"slug"`
		ContractAddresses   []interface{} `json:"contract_addresses"`
		InternalTempAgoraID string        `json:"_internal_temp_agora_id"`
		Parameters          struct {
			AssetKey        string    `json:"asset_key"`
			AssetID         string    `json:"asset_id"`
			Start           time.Time `json:"start"`
			End             time.Time `json:"end"`
			Interval        string    `json:"interval"`
			Order           string    `json:"order"`
			Format          string    `json:"format"`
			TimestampFormat string    `json:"timestamp_format"`
			Columns         []string  `json:"columns"`
		} `json:"parameters"`
		Schema struct {
			MetricID     string `json:"metric_id"`
			Name         string `json:"name"`
			Description  string `json:"description"`
			ValuesSchema struct {
				Timestamp string `json:"timestamp"`
				Open      string `json:"open"`
				High      string `json:"high"`
				Low       string `json:"low"`
				Close     string `json:"close"`
				Volume    string `json:"volume"`
			} `json:"values_schema"`
			MinimumInterval   string `json:"minimum_interval"`
			SourceAttribution []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"source_attribution"`
		} `json:"schema"`
		Values [][]float64 `json:"values"`
	} `json:"data"`
}
