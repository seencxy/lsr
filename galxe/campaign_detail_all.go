package galxe

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Campaign详情请求参数
type CampaignDetailAllRequest struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		Address     string `json:"address"`
		WithAddress bool   `json:"withAddress"`
		Id          string `json:"id"`
	} `json:"variables"`
	Query string `json:"query"`
}

// Campaign详情响应参数
type CampaignDetailAllResponse struct {
	Data struct {
		Campaign struct {
			CoHostSpaces []interface{} `json:"coHostSpaces"`
			BannerUrl    string        `json:"bannerUrl"`
			Id           string        `json:"id"`
			Thumbnail    string        `json:"thumbnail"`
			RewardName   string        `json:"rewardName"`
			Type         string        `json:"type"`
			Gamification struct {
				Id          string        `json:"id"`
				Type        string        `json:"type"`
				Typename    string        `json:"__typename"`
				ForgeConfig interface{}   `json:"forgeConfig"`
				Nfts        []interface{} `json:"nfts"`
				Airdrop     interface{}   `json:"airdrop"`
			} `json:"gamification"`
			Typename              string      `json:"__typename"`
			NumberID              int         `json:"numberID"`
			Chain                 string      `json:"chain"`
			SpaceStation          interface{} `json:"spaceStation"`
			Name                  string      `json:"name"`
			Cap                   int         `json:"cap"`
			Info                  string      `json:"info"`
			UseCred               bool        `json:"useCred"`
			Formula               string      `json:"formula"`
			Status                string      `json:"status"`
			Creator               string      `json:"creator"`
			GasType               string      `json:"gasType"`
			IsPrivate             bool        `json:"isPrivate"`
			CreatedAt             string      `json:"createdAt"`
			RequirementInfo       string      `json:"requirementInfo"`
			Description           string      `json:"description"`
			EnableWhitelist       bool        `json:"enableWhitelist"`
			StartTime             int         `json:"startTime"`
			EndTime               int         `json:"endTime"`
			RequireEmail          bool        `json:"requireEmail"`
			RequireUsername       bool        `json:"requireUsername"`
			BlacklistCountryCodes string      `json:"blacklistCountryCodes"`
			WhitelistRegions      string      `json:"whitelistRegions"`
			RewardType            string      `json:"rewardType"`
			DistributionType      string      `json:"distributionType"`
			ClaimEndTime          interface{} `json:"claimEndTime"`
			LoyaltyPoints         int         `json:"loyaltyPoints"`
			TokenRewardContract   interface{} `json:"tokenRewardContract"`
			TokenReward           struct {
				UserTokenAmount      string `json:"userTokenAmount"`
				TokenAddress         string `json:"tokenAddress"`
				DepositedTokenAmount string `json:"depositedTokenAmount"`
				TokenRewardId        int    `json:"tokenRewardId"`
				TokenDecimal         string `json:"tokenDecimal"`
				TokenLogo            string `json:"tokenLogo"`
				TokenSymbol          string `json:"tokenSymbol"`
				Typename             string `json:"__typename"`
			} `json:"tokenReward"`
			NftHolderSnapshot struct {
				HolderSnapshotBlock int    `json:"holderSnapshotBlock"`
				Typename            string `json:"__typename"`
			} `json:"nftHolderSnapshot"`
			WhitelistInfo struct {
				Address                           string `json:"address"`
				MaxCount                          int    `json:"maxCount"`
				UsedCount                         int    `json:"usedCount"`
				ClaimedLoyaltyPoints              int    `json:"claimedLoyaltyPoints"`
				CurrentPeriodClaimedLoyaltyPoints int    `json:"currentPeriodClaimedLoyaltyPoints"`
				CurrentPeriodMaxLoyaltyPoints     int    `json:"currentPeriodMaxLoyaltyPoints"`
				Typename                          string `json:"__typename"`
			} `json:"whitelistInfo"`
			WhitelistSubgraph interface{}   `json:"whitelistSubgraph"`
			Creds             []interface{} `json:"creds"`
			CredentialGroups  []struct {
				Id          string `json:"id"`
				Description string `json:"description"`
				Credentials []struct {
					Id                    string      `json:"id"`
					Name                  string      `json:"name"`
					Type                  string      `json:"type"`
					CredType              string      `json:"credType"`
					CredSource            string      `json:"credSource"`
					ReferenceLink         string      `json:"referenceLink"`
					Description           string      `json:"description"`
					LastUpdate            int         `json:"lastUpdate"`
					SyncStatus            string      `json:"syncStatus"`
					CredContractNFTHolder interface{} `json:"credContractNFTHolder"`
					Chain                 string      `json:"chain"`
					Eligible              int         `json:"eligible"`
					Subgraph              interface{} `json:"subgraph"`
					Metadata              struct {
						VisitLink          interface{} `json:"visitLink"`
						GitcoinPassport    interface{} `json:"gitcoinPassport"`
						CampaignReferral   interface{} `json:"campaignReferral"`
						GalxeScore         interface{} `json:"galxeScore"`
						RestApi            interface{} `json:"restApi"`
						WalletBalance      interface{} `json:"walletBalance"`
						LensProfileFollow  interface{} `json:"lensProfileFollow"`
						Graphql            interface{} `json:"graphql"`
						LensPostUpvote     interface{} `json:"lensPostUpvote"`
						LensPostMirror     interface{} `json:"lensPostMirror"`
						MultiDimensionRest interface{} `json:"multiDimensionRest"`
						Typename           string      `json:"__typename"`
					} `json:"metadata"`
					DimensionConfig string      `json:"dimensionConfig"`
					Value           interface{} `json:"value"`
					Typename        string      `json:"__typename"`
				} `json:"credentials"`
				ConditionRelation string `json:"conditionRelation"`
				Conditions        []struct {
					Expression      string   `json:"expression"`
					Eligible        bool     `json:"eligible"`
					EligibleAddress []string `json:"eligibleAddress"`
					Typename        string   `json:"__typename"`
				} `json:"conditions"`
				Rewards []struct {
					Expression  string `json:"expression"`
					Eligible    bool   `json:"eligible"`
					RewardCount int    `json:"rewardCount"`
					RewardType  string `json:"rewardType"`
					Typename    string `json:"__typename"`
				} `json:"rewards"`
				RewardAttrVals       interface{} `json:"rewardAttrVals"`
				ClaimedLoyaltyPoints int         `json:"claimedLoyaltyPoints"`
				Typename             string      `json:"__typename"`
			} `json:"credentialGroups"`
			Dao struct {
				Id         string `json:"id"`
				Name       string `json:"name"`
				Logo       string `json:"logo"`
				Alias      string `json:"alias"`
				IsVerified bool   `json:"isVerified"`
				Typename   string `json:"__typename"`
				NftCores   struct {
					List     []interface{} `json:"list"`
					Typename string        `json:"__typename"`
				} `json:"nftCores"`
			} `json:"dao"`
			RewardInfo struct {
				DiscordRole             interface{} `json:"discordRole"`
				Premint                 interface{} `json:"premint"`
				LoyaltyPoints           interface{} `json:"loyaltyPoints"`
				LoyaltyPointsMysteryBox interface{} `json:"loyaltyPointsMysteryBox"`
				Typename                string      `json:"__typename"`
			} `json:"rewardInfo"`
			Participants struct {
				ParticipantsCount  int    `json:"participantsCount"`
				BountyWinnersCount int    `json:"bountyWinnersCount"`
				Typename           string `json:"__typename"`
			} `json:"participants"`
			TaskConfig          interface{} `json:"taskConfig"`
			ReferralCode        interface{} `json:"referralCode"`
			RecurringType       string      `json:"recurringType"`
			LatestRecurringTime interface{} `json:"latestRecurringTime"`
			UserParticipants    struct {
				List     []interface{} `json:"list"`
				Typename string        `json:"__typename"`
			} `json:"userParticipants"`
			Space struct {
				Id             string   `json:"id"`
				Name           string   `json:"name"`
				Info           string   `json:"info"`
				Thumbnail      string   `json:"thumbnail"`
				Alias          string   `json:"alias"`
				Links          string   `json:"links"`
				IsVerified     bool     `json:"isVerified"`
				DiscordGuildID string   `json:"discordGuildID"`
				FollowersCount int      `json:"followersCount"`
				Typename       string   `json:"__typename"`
				IsAdmin        bool     `json:"isAdmin"`
				IsFollowing    bool     `json:"isFollowing"`
				Categories     []string `json:"categories"`
			} `json:"space"`
			IsBookmarked         bool        `json:"isBookmarked"`
			ClaimedLoyaltyPoints int         `json:"claimedLoyaltyPoints"`
			ParentCampaign       interface{} `json:"parentCampaign"`
			IsSequencial         bool        `json:"isSequencial"`
			NumNFTMinted         int         `json:"numNFTMinted"`
			ChildrenCampaigns    interface{} `json:"childrenCampaigns"`
		} `json:"campaign"`
	} `json:"data"`
}

// 实现获取Campaign详情
func CampaignDetailAll(client http.Client, auth string, address string, id string) (CampaignDetailAllResponse, error) {
	data := CampaignDetailAllRequest{
		OperationName: "CampaignDetailAll",
		Query:         "query CampaignDetailAll($id: ID!, $address: String!, $withAddress: Boolean!) {\n  campaign(id: $id) {\n    coHostSpaces {\n      ...SpaceDetail\n      isAdmin(address: $address) @include(if: $withAddress)\n      isFollowing @include(if: $withAddress)\n      followersCount\n      categories\n      __typename\n    }\n    bannerUrl\n    ...CampaignDetailFrag\n    userParticipants(address: $address, first: 1) @include(if: $withAddress) {\n      list {\n        status\n        premintTo\n        __typename\n      }\n      __typename\n    }\n    space {\n      ...SpaceDetail\n      isAdmin(address: $address) @include(if: $withAddress)\n      isFollowing @include(if: $withAddress)\n      followersCount\n      categories\n      __typename\n    }\n    isBookmarked(address: $address) @include(if: $withAddress)\n    claimedLoyaltyPoints(address: $address) @include(if: $withAddress)\n    parentCampaign {\n      id\n      isSequencial\n      thumbnail\n      __typename\n    }\n    isSequencial\n    numNFTMinted\n    childrenCampaigns {\n      space {\n        ...SpaceDetail\n        isAdmin(address: $address) @include(if: $withAddress)\n        isFollowing @include(if: $withAddress)\n        followersCount\n        categories\n        __typename\n      }\n      ...CampaignDetailFrag\n      claimedLoyaltyPoints(address: $address) @include(if: $withAddress)\n      userParticipants(address: $address, first: 1) @include(if: $withAddress) {\n        list {\n          status\n          __typename\n        }\n        __typename\n      }\n      parentCampaign {\n        id\n        isSequencial\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CampaignDetailFrag on Campaign {\n  id\n  ...CampaignMedia\n  ...CampaignForgePage\n  name\n  numberID\n  type\n  cap\n  info\n  useCred\n  formula\n  status\n  creator\n  thumbnail\n  gasType\n  isPrivate\n  createdAt\n  requirementInfo\n  description\n  enableWhitelist\n  chain\n  startTime\n  endTime\n  requireEmail\n  requireUsername\n  blacklistCountryCodes\n  whitelistRegions\n  rewardType\n  distributionType\n  rewardName\n  claimEndTime\n  loyaltyPoints\n  tokenRewardContract {\n    id\n    address\n    chain\n    __typename\n  }\n  tokenReward {\n    userTokenAmount\n    tokenAddress\n    depositedTokenAmount\n    tokenRewardId\n    tokenDecimal\n    tokenLogo\n    tokenSymbol\n    __typename\n  }\n  nftHolderSnapshot {\n    holderSnapshotBlock\n    __typename\n  }\n  spaceStation {\n    id\n    address\n    chain\n    __typename\n  }\n  ...WhitelistInfoFrag\n  ...WhitelistSubgraphFrag\n  gamification {\n    ...GamificationDetailFrag\n    __typename\n  }\n  creds {\n    ...CredForAddress\n    __typename\n  }\n  credentialGroups(address: $address) {\n    ...CredentialGroupForAddress\n    __typename\n  }\n  dao {\n    ...DaoSnap\n    nftCores {\n      list {\n        capable\n        marketLink\n        contractAddress\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  rewardInfo {\n    discordRole {\n      guildId\n      guildName\n      roleId\n      roleName\n      inviteLink\n      __typename\n    }\n    premint {\n      startTime\n      endTime\n      chain\n      price\n      totalSupply\n      contractAddress\n      banner\n      __typename\n    }\n    loyaltyPoints {\n      points\n      __typename\n    }\n    loyaltyPointsMysteryBox {\n      points\n      weight\n      __typename\n    }\n    __typename\n  }\n  participants {\n    participantsCount\n    bountyWinnersCount\n    __typename\n  }\n  taskConfig(address: $address) {\n    participateCondition {\n      conditions {\n        ...ExpressionEntity\n        __typename\n      }\n      conditionalFormula\n      eligible\n      __typename\n    }\n    rewardConfigs {\n      conditions {\n        ...ExpressionEntity\n        __typename\n      }\n      conditionalFormula\n      description\n      rewards {\n        ...ExpressionReward\n        __typename\n      }\n      eligible\n      rewardAttrVals {\n        attrName\n        attrTitle\n        attrVal\n        __typename\n      }\n      __typename\n    }\n    referralConfig {\n      conditions {\n        ...ExpressionEntity\n        __typename\n      }\n      conditionalFormula\n      description\n      rewards {\n        ...ExpressionReward\n        __typename\n      }\n      eligible\n      rewardAttrVals {\n        attrName\n        attrTitle\n        attrVal\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  referralCode(address: $address)\n  recurringType\n  latestRecurringTime\n  __typename\n}\n\nfragment DaoSnap on DAO {\n  id\n  name\n  logo\n  alias\n  isVerified\n  __typename\n}\n\nfragment CampaignMedia on Campaign {\n  thumbnail\n  rewardName\n  type\n  gamification {\n    id\n    type\n    __typename\n  }\n  __typename\n}\n\nfragment CredForAddress on Cred {\n  id\n  name\n  type\n  credType\n  credSource\n  referenceLink\n  description\n  lastUpdate\n  syncStatus\n  credContractNFTHolder {\n    timestamp\n    __typename\n  }\n  chain\n  eligible(address: $address)\n  subgraph {\n    endpoint\n    query\n    expression\n    __typename\n  }\n  metadata {\n    ...CredMetaData\n    __typename\n  }\n  dimensionConfig\n  value {\n    gitcoinPassport {\n      score\n      lastScoreTimestamp\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment CredMetaData on CredMetadata {\n  visitLink {\n    link\n    __typename\n  }\n  gitcoinPassport {\n    score {\n      title\n      type\n      description\n      config\n      __typename\n    }\n    lastScoreTimestamp {\n      title\n      type\n      description\n      config\n      __typename\n    }\n    __typename\n  }\n  campaignReferral {\n    count {\n      title\n      type\n      description\n      config\n      __typename\n    }\n    __typename\n  }\n  galxeScore {\n    dimensions {\n      id\n      type\n      title\n      description\n      config\n      values {\n        name\n        type\n        value\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  restApi {\n    url\n    method\n    headers {\n      key\n      value\n      __typename\n    }\n    postBody\n    expression\n    __typename\n  }\n  walletBalance {\n    contractAddress\n    snapshotTimestamp\n    chain\n    balance {\n      type\n      title\n      description\n      config\n      __typename\n    }\n    LastSyncBlock\n    LastSyncTimestamp\n    __typename\n  }\n  lensProfileFollow {\n    handle\n    __typename\n  }\n  graphql {\n    url\n    query\n    expression\n    __typename\n  }\n  lensPostUpvote {\n    postId\n    __typename\n  }\n  lensPostMirror {\n    postId\n    __typename\n  }\n  multiDimensionRest {\n    url\n    method\n    headers {\n      key\n      value\n      __typename\n    }\n    postBody\n    expression\n    dimensions {\n      id\n      type\n      title\n      description\n      config\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment CredentialGroupForAddress on CredentialGroup {\n  id\n  description\n  credentials {\n    ...CredForAddress\n    __typename\n  }\n  conditionRelation\n  conditions {\n    expression\n    eligible\n    ...CredentialGroupConditionForVerifyButton\n    __typename\n  }\n  rewards {\n    expression\n    eligible\n    rewardCount\n    rewardType\n    __typename\n  }\n  rewardAttrVals {\n    attrName\n    attrTitle\n    attrVal\n    __typename\n  }\n  claimedLoyaltyPoints\n  __typename\n}\n\nfragment CredentialGroupConditionForVerifyButton on CredentialGroupCondition {\n  expression\n  eligibleAddress\n  __typename\n}\n\nfragment WhitelistInfoFrag on Campaign {\n  id\n  whitelistInfo(address: $address) {\n    address\n    maxCount\n    usedCount\n    claimedLoyaltyPoints\n    currentPeriodClaimedLoyaltyPoints\n    currentPeriodMaxLoyaltyPoints\n    __typename\n  }\n  __typename\n}\n\nfragment WhitelistSubgraphFrag on Campaign {\n  id\n  whitelistSubgraph {\n    query\n    endpoint\n    expression\n    variable\n    __typename\n  }\n  __typename\n}\n\nfragment GamificationDetailFrag on Gamification {\n  id\n  type\n  nfts {\n    nft {\n      id\n      animationURL\n      category\n      powah\n      image\n      name\n      treasureBack\n      nftCore {\n        ...NftCoreInfoFrag\n        __typename\n      }\n      traits {\n        name\n        value\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  airdrop {\n    name\n    contractAddress\n    token {\n      address\n      icon\n      symbol\n      __typename\n    }\n    merkleTreeUrl\n    addressInfo(address: $address) {\n      index\n      amount {\n        amount\n        ether\n        __typename\n      }\n      proofs\n      __typename\n    }\n    __typename\n  }\n  forgeConfig {\n    minNFTCount\n    maxNFTCount\n    requiredNFTs {\n      nft {\n        category\n        powah\n        image\n        name\n        nftCore {\n          capable\n          contractAddress\n          __typename\n        }\n        __typename\n      }\n      count\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment NftCoreInfoFrag on NFTCore {\n  id\n  capable\n  chain\n  contractAddress\n  name\n  symbol\n  dao {\n    id\n    name\n    logo\n    alias\n    __typename\n  }\n  __typename\n}\n\nfragment ExpressionEntity on ExprEntity {\n  cred {\n    id\n    name\n    type\n    credType\n    credSource\n    referenceLink\n    description\n    lastUpdate\n    chain\n    eligible(address: $address)\n    metadata {\n      visitLink {\n        link\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  attrs {\n    attrName\n    operatorSymbol\n    targetValue\n    __typename\n  }\n  attrFormula\n  eligible\n  __typename\n}\n\nfragment ExpressionReward on ExprReward {\n  arithmetics {\n    ...ExpressionEntity\n    __typename\n  }\n  arithmeticFormula\n  rewardType\n  rewardCount\n  rewardVal\n  __typename\n}\n\nfragment CampaignForgePage on Campaign {\n  id\n  numberID\n  chain\n  spaceStation {\n    address\n    __typename\n  }\n  gamification {\n    forgeConfig {\n      maxNFTCount\n      minNFTCount\n      requiredNFTs {\n        nft {\n          category\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment SpaceDetail on Space {\n  id\n  name\n  info\n  thumbnail\n  alias\n  links\n  isVerified\n  discordGuildID\n  followersCount\n  __typename\n}\n",
		Variables: struct {
			Address     string `json:"address"`
			WithAddress bool   `json:"withAddress"`
			Id          string `json:"id"`
		}{
			Address:     address,
			WithAddress: true,
			Id:          id,
		},
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		return *new(CampaignDetailAllResponse), err
	}

	// 创建请求
	request, err := http.NewRequest("POST", "https://graphigo.prd.galaxy.eco/query", bytes.NewReader(marshal))
	if err != nil {
		return *new(CampaignDetailAllResponse), err
	}

	// 设置头部
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Origin", "https://galxe.com")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	request.Header.Add("Sec-Ch-Ua", "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"Google Chrome\";v=\"116\"")
	request.Header.Add("Authorization", auth)

	all, err := client.Do(request)
	if err != nil {
		return *new(CampaignDetailAllResponse), err
	}

	if all.StatusCode != 200 {
		return *new(CampaignDetailAllResponse), err
	}

	readAll, err := io.ReadAll(all.Body)
	if err != nil {
		return *new(CampaignDetailAllResponse), err
	}

	var body CampaignDetailAllResponse

	err = json.Unmarshal(readAll, &body)
	if err != nil {
		return *new(CampaignDetailAllResponse), err
	}

	return body, nil
}
