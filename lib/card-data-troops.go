package lib

var (
	// Common
	KNIGHT = Card{
		NAME:   "Knight",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `A tough melee fighter. The Barbarian's handsome, cultured cousin. Rumor has it that he was knighted based on the sheer awesomeness of his mustache alone.`,
		COST:   3,
		HP:     V600[0:12:12],
		DPS:    []int{68, 74, 81, 90, 99, 109, 120, 130, 144, 158, 174, 190},
		DAM:    V75D[0:12:12],
		HSPD:   1.1,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    MELEE,
		DTIME:  1,
	}
	ARCHERS = Card{
		NAME:   "Archers",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `A pair of unarmored ranged attackers. They'll help you with ground and air unit attacks, but you're on your own with coloring your hair.`,
		COST:   3,
		HP:     V125[0:12:12],
		DPS:    []int{33, 36, 40, 44, 48, 53, 58, 64, 70, 77, 85, 93},
		DAM:    V40D[0:12:12],
		HSPD:   1.2,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    5.5,
		DTIME:  1,
		COUNT:  2,
	}
	BOMBER = Card{
		NAME:   "Bomber",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `Small, lightly protected skeleton that throws bombs. Deals area damage that can wipe out a swarm of enemies.`,
		COST:   3,
		HP:     V150[0:12:12],
		DPS:    []int{52, 57, 63, 70, 76, 84, 92, 101, 111, 122, 134, 147},
		ADAM:   V100D[0:12:12],
		HSPD:   1.9,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    5,
		DTIME:  1,
	}
	GOBLINS = Card{
		NAME:   "Goblins",
		ARENA:  ARENA_1,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `Three fast, unarmored melee attackers. Small, fast, green and mean!`,
		COST:   2,
		HP:     V80[0:12:12],
		DPS:    []int{45, 50, 54, 60, 66, 72, 80, X, 96, 105, 116, 127},
		DAM:    V50D[0:12:12],
		HSPD:   1.1,
		TGTS:   GROUND,
		SPD:    VERY_FAST,
		RNG:    MELEE,
		DTIME:  1,
		COUNT:  3,
	}
	SPEAR_GOBLINS = Card{
		NAME:   "Spear Goblins",
		ARENA:  ARENA_1,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `Three unarmored ranged attackers. Who the heck taught these guys to throw spears!?! Who thought that was a good idea?!`,
		COST:   2,
		HP:     V52[0:12:12],
		DPS:    []int{18, 20, 22, 23, 26, 29, 32, 35, 38, 42, 46, 51},
		DAM:    V24D[0:12:12],
		HSPD:   1.3,
		TGTS:   AIR_AND_GROUND,
		SPD:    VERY_FAST,
		RNG:    5.5,
		DTIME:  1,
		COUNT:  3,
	}
	SKELETONS = Card{
		NAME:   "Skeletons",
		ARENA:  ARENA_2,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `Four fast, very weak melee fighters. Swarm your enemies with this pile of bones!`,
		COST:   1,
		HP:     V30[0:12:12],
		DPS:    V30[0:12:12],
		DAM:    V30[0:12:12],
		HSPD:   1,
		TGTS:   GROUND,
		SPD:    FAST,
		RNG:    MELEE,
		DTIME:  1,
		COUNT:  4,
	}
	MINIONS = Card{
		NAME:   "Minions",
		ARENA:  ARENA_2,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `Three fast, unarmored flying attackers. Roses are red, minions are blue, they can fly, and will crush you!`,
		COST:   3,
		HP:     V90[0:12:12],
		DPS:    V40D[0:12:12],
		DAM:    V40D[0:12:12],
		HSPD:   1,
		TGTS:   AIR_AND_GROUND,
		SPD:    FAST,
		RNG:    2.5,
		DTIME:  1,
		COUNT:  3,
	}
	BARBARIANS = Card{
		NAME:   "Barbarians",
		ARENA:  ARENA_3,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `A horde of melee attackers with mean mustaches and even meaner tempers.`,
		COST:   5,
		HP:     V300[0:12:12],
		DPS:    []int{50, X, X, 66, 72, 80, 88, 96, 106, 116, 128, 140},
		DAM:    V75D[0:12:12],
		HSPD:   1.5,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    MELEE,
		DTIME:  1,
		COUNT:  4,
	}
	MINION_HORDE = Card{
		NAME:   "Minion Horde",
		ARENA:  ARENA_4,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `Six fast, unarmored flying attackers. Three's a crowd, six is a horde!`,
		COST:   5,
		HP:     V90[0:12:12],
		DPS:    V40D[0:12:12],
		DAM:    V40D[0:12:12],
		HSPD:   1,
		TGTS:   AIR_AND_GROUND,
		SPD:    FAST,
		RNG:    2.5,
		DTIME:  1,
		COUNT:  6,
	}

	// Rare
	GIANT = Card{
		NAME:   "Giant",
		ARENA:  ARENA_0,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   `Slow but durable, only attacks buildings. A real one-man wrecking crew!`,
		COST:   5,
		HP:     V2000[0:10:10],
		DPS:    []int{84, X, 101, 111, 122, 134, X, X, X, 195},
		DAM:    V126D[0:10:10],
		HSPD:   1.5,
		TGTS:   BUILDINGS,
		SPD:    SLOW,
		RNG:    MELEE,
		DTIME:  1,
	}
	MUSKETEER = Card{
		NAME:   "Musketeer",
		ARENA:  ARENA_0,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   `Don't be fooled by her delicately coiffed hair, the Musketeer is a mean shot with her trusty boomstick.`,
		COST:   4,
		HP:     V340[0:10:10],
		DPS:    []int{90, X, 110, 120, 132, X, X, 175, 192, 211},
		DAM:    V100D[0:10:10],
		HSPD:   1.1,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    6.5,
		DTIME:  1,
	}
	MINI_PEKKA = Card{
		NAME:   "Mini P.E.K.K.A",
		ARENA:  ARENA_0,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   `The arena is a certified butterfly-free zone. No distractions for P.E.K.K.A, only destruction.`,
		COST:   4,
		HP:     V600[0:10:10],
		DPS:    []int{180, 198, 218, 240, 263, X, 317, 348, 382, 420},
		DAM:    V325D[0:10:10],
		HSPD:   1.8,
		TGTS:   GROUND,
		SPD:    FAST,
		RNG:    MELEE,
		DTIME:  1,
	}
	VALKYRIE = Card{
		NAME:   "Valkyrie",
		ARENA:  ARENA_1,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   `Tough melee fighter, deals area damage around her. Swarm or horde, no problem! She can take them all out with a few spins.`,
		COST:   4,
		HP:     V880[0:10:10],
		DPS:    []int{80, X, X, 106, 116, X, X, X, X, 186},
		DAM:    V120D[0:10:10],
		HSPD:   1.5,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    MELEE,
		DTIME:  1,
	}
	HOG_RIDER = Card{
		NAME:   "Hog Rider",
		ARENA:  ARENA_4,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   `Fast melee troop that targets buildings and can jump over the river. He followed the echoing call of "Hog Riderrrrr" all the way through the arena doors.`,
		COST:   4,
		HP:     V800[0:10:10],
		DPS:    []int{106, 117, 128, 141, 155, X, X, 205, 226, 248},
		DAM:    V160D[0:10:10],
		HSPD:   1.5,
		TGTS:   BUILDINGS,
		SPD:    VERY_FAST,
		RNG:    MELEE,
		DTIME:  1,
	}
	WIZARD = Card{
		NAME:   "Wizard",
		ARENA:  ARENA_5,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   `The most awesome man to ever set foot in the arena, the Wizard will blow you away with his handsomeness... and/or fireballs.`,
		COST:   5,
		HP:     V340[0:10:10],
		DPS:    []int{76, X, 92, 101, 111, X, X, 147, 161, 177},
		ADAM:   V130D[0:10:10],
		HSPD:   1.7,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    5.5,
		DTIME:  1,
	}

	// Epic
	WITCH = Card{
		NAME:   "Witch",
		ARENA:  ARENA_0,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   `Summons skeletons, shoots destructo beams, has glowing pink eyes that unfortunately don't shoot lasers.`,
		COST:   5,
		HP:     V500[0:8:8],
		DPS:    []int{54, 58, 64, X, X, X, X, 104},
		ADAM:   V38D[0:8:8],
		SKE_LV: []int{6, 7, 8, 9, 10, 11, 12, 13},
		SSPD:   7.5,
		HSPD:   0.7,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    5.5,
		DTIME:  1,
	}
	SKELETON_ARMY = Card{
		NAME:   "Skeleton Army",
		ARENA:  ARENA_0,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   `Summons an army of skeletons. Meet Larry and his friends Harry, Gerry, Terry, Mary, etc.`,
		COST:   4,
		HP:     V30[0:8:8],
		DPS:    V30[0:8:8],
		DAM:    V30[0:8:8],
		HSPD:   1,
		TGTS:   GROUND,
		SPD:    FAST,
		RNG:    MELEE,
		DTIME:  1,
		COUNT:  20,
	}
	BABY_DRAGON = Card{
		NAME:   "Baby Dragon",
		ARENA:  ARENA_0,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   `Flying troop that deals area damage. Baby dragons hatch cute, hungry, and ready for a barbeque.`,
		COST:   4,
		HP:     V800[0:8:8],
		DPS:    []int{55, 61, 67, 73, 81, 88, 97, 107},
		ADAM:   V100D[0:8:8],
		HSPD:   1.8,
		TGTS:   AIR_AND_GROUND,
		SPD:    FAST,
		RNG:    3.5,
		DTIME:  1,
	}
	PRINCE = Card{
		NAME:   "Prince",
		ARENA:  ARENA_0,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   `Don't let the little pony fool you. Once the Prince gets a running start, you WILL be trampled. Does 2x damage once he gets charging.`,
		COST:   5,
		HP:     V1100[0:8:8],
		DPS:    []int{146, 161, 177, 194, 214, 234, 258, 282},
		DAM:    V220D[0:8:8],
		HSPD:   1.5,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    2.5,
		DTIME:  1,
	}
	GIANT_SKELETON = Card{
		NAME:   "Giant Skeleton",
		ARENA:  ARENA_2,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   `The bigger the skeleton, the bigger the bomb. Carries a bomb that blows up when the Giant Skeleton dies.`,
		COST:   6,
		HP:     V2000[0:8:8],
		DPS:    []int{66, 73, 80, X, X, 106, 117, 128},
		DAM:    V100D[0:8:8],
		DDAM:   V720D[0:8:8],
		HSPD:   1.5,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    MELEE,
		DTIME:  1,
	}
	BALLOON = Card{
		NAME:   "Balloon",
		ARENA:  ARENA_2,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   `As pretty as they are, you won't want a parade of THESE balloons showing up on the horizon. Drops powerful bombs and when shot down, crashes dealing area damage.`,
		COST:   5,
		HP:     V1050[0:8:8],
		DPS:    []int{200, 220, X, X, 292, X, 352, X},
		DAM:    V600[0:8:8],
		DDAM:   V100D[0:8:8],
		HSPD:   3,
		TGTS:   BUILDINGS,
		SPD:    MEDIUM,
		RNG:    MELEE,
		DTIME:  1,
	}
	PEKKA = Card{
		NAME:   "P.E.K.K.A",
		ARENA:  ARENA_4,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   `A heavily armored, slow melee fighter. Swings from the hip but packs a huge punch.`,
		COST:   7,
		HP:     V2600[0:8:8],
		DPS:    []int{250, 275, X, X, X, X, 440, 482},
		DAM:    V450[0:8:8],
		HSPD:   1.8,
		TGTS:   GROUND,
		SPD:    SLOW,
		RNG:    MELEE,
		DTIME:  3,
	}
	GOLEM = Card{
		NAME:   "Golem",
		ARENA:  ARENA_6,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   `Slow but durable, only attacks buildings. When destroyed, explosively splits into two Golemites and deals area damage!`,
		COST:   8,
		HP:     V3000[0:8:8],
		DPS:    []int{74, X, X, X, 108, 118, 130, 143},
		DAM:    V186D[0:8:8],
		DDAM:   V186D[0:8:8],
		HSPD:   2.5,
		TGTS:   BUILDINGS,
		SPD:    SLOW,
		RNG:    MELEE,
		DTIME:  3,
	}
)
