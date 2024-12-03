package main

func dayTwoResult() int {
	reports := [][]int{
		{3, 6, 7, 9, 11, 8},
		{21, 24, 25, 26, 29, 30, 32, 32},
		{29, 32, 33, 34, 35, 37, 38, 42},
		{54, 55, 57, 58, 60, 61, 63, 70},
		{59, 61, 60, 63, 65, 68, 71},
		{43, 44, 45, 44, 46, 44},
		{73, 75, 73, 74, 75, 75},
		{82, 85, 82, 85, 89},
		{79, 82, 84, 83, 85, 92},
		{91, 92, 93, 93, 95},
		{30, 33, 35, 37, 37, 39, 40, 38},
		{74, 75, 76, 78, 78, 81, 81},
		{25, 26, 29, 29, 30, 33, 35, 39},
		{71, 72, 72, 74, 81},
		{48, 51, 55, 57, 59, 60, 61, 63},
		{57, 59, 61, 62, 64, 65, 69, 66},
		{86, 88, 90, 92, 95, 99, 99},
		{2, 3, 6, 7, 9, 13, 17},
		{84, 85, 89, 92, 98},
		{1, 2, 8, 9, 11, 12},
		{36, 39, 46, 47, 49, 51, 53, 50},
		{59, 60, 63, 70, 72, 73, 75, 75},
		{6, 9, 12, 18, 19, 21, 22, 26},
		{69, 70, 77, 80, 81, 84, 89},
		{20, 17, 18, 21, 23, 25},
		{17, 14, 15, 16, 17, 20, 19},
		{69, 67, 68, 71, 71},
		{27, 24, 27, 30, 31, 33, 37},
		{20, 19, 20, 22, 28},
		{51, 48, 50, 52, 51, 53, 56, 58},
		{6, 5, 7, 8, 10, 8, 7},
		{74, 72, 74, 71, 71},
		{36, 34, 36, 37, 34, 35, 39},
		{6, 3, 5, 8, 6, 8, 11, 16},
		{76, 74, 76, 79, 82, 82, 84},
		{38, 36, 38, 39, 39, 42, 43, 40},
		{33, 31, 32, 35, 37, 37, 37},
		{78, 75, 75, 78, 82},
		{67, 66, 67, 69, 69, 71, 77},
		{86, 84, 85, 87, 91, 93},
		{55, 52, 56, 58, 60, 58},
		{6, 5, 8, 11, 15, 17, 19, 19},
		{65, 62, 66, 68, 71, 75},
		{61, 60, 62, 66, 67, 70, 75},
		{3, 1, 2, 3, 10, 12},
		{83, 81, 84, 86, 92, 90},
		{45, 43, 45, 51, 52, 52},
		{58, 55, 61, 62, 66},
		{55, 53, 56, 58, 60, 62, 68, 75},
		{65, 65, 66, 68, 69, 72, 75},
		{75, 75, 76, 79, 82, 81},
		{50, 50, 52, 54, 57, 60, 63, 63},
		{11, 11, 14, 16, 17, 19, 23},
		{56, 56, 58, 59, 61, 64, 65, 72},
		{87, 87, 89, 92, 90, 93},
		{73, 73, 74, 76, 79, 78, 81, 80},
		{56, 56, 59, 58, 60, 63, 66, 66},
		{16, 16, 19, 17, 18, 20, 24},
		{70, 70, 71, 73, 74, 77, 76, 83},
		{71, 71, 73, 73, 74, 75, 77, 80},
		{30, 30, 31, 33, 33, 32},
		{73, 73, 76, 77, 79, 81, 81, 81},
		{42, 42, 44, 46, 48, 48, 52},
		{53, 53, 53, 56, 61},
		{62, 62, 64, 65, 66, 70, 72, 75},
		{47, 47, 50, 52, 55, 57, 61, 60},
		{11, 11, 13, 14, 17, 21, 24, 24},
		{33, 33, 37, 38, 39, 41, 45},
		{46, 46, 48, 50, 54, 57, 63},
		{51, 51, 53, 58, 60, 63},
		{10, 10, 11, 14, 15, 16, 21, 18},
		{4, 4, 6, 9, 14, 16, 16},
		{54, 54, 60, 61, 65},
		{36, 36, 39, 45, 48, 53},
		{58, 62, 65, 67, 69, 70, 71, 73},
		{64, 68, 69, 70, 72, 75, 74},
		{89, 93, 96, 97, 98, 99, 99},
		{58, 62, 65, 66, 69, 70, 74},
		{13, 17, 19, 22, 25, 27, 34},
		{51, 55, 54, 57, 59, 61},
		{69, 73, 76, 78, 80, 79, 81, 80},
		{30, 34, 37, 34, 34},
		{5, 9, 10, 7, 10, 14},
		{74, 78, 79, 76, 77, 83},
		{87, 91, 94, 97, 98, 98, 99},
		{67, 71, 73, 75, 75, 74},
		{68, 72, 72, 75, 75},
		{77, 81, 81, 82, 86},
		{16, 20, 21, 22, 22, 27},
		{31, 35, 39, 40, 43},
		{85, 89, 93, 94, 95, 97, 99, 98},
		{54, 58, 62, 63, 65, 68, 69, 69},
		{78, 82, 83, 87, 90, 91, 94, 98},
		{17, 21, 25, 27, 28, 29, 34},
		{34, 38, 39, 45, 48, 50, 52, 54},
		{32, 36, 38, 40, 45, 44},
		{65, 69, 74, 76, 76},
		{12, 16, 17, 23, 25, 29},
		{26, 30, 33, 35, 41, 47},
		{3, 8, 9, 11, 14, 17},
		{56, 63, 64, 67, 64},
		{52, 57, 58, 59, 61, 61},
		{77, 83, 85, 87, 89, 90, 94},
		{7, 13, 15, 17, 19, 21, 22, 27},
		{52, 58, 59, 56, 57, 60, 63},
		{16, 22, 24, 26, 23, 24, 27, 24},
		{66, 73, 75, 76, 73, 76, 76},
		{9, 15, 14, 17, 18, 19, 22, 26},
		{17, 23, 25, 26, 28, 31, 30, 35},
		{82, 87, 88, 90, 90, 92, 93},
		{63, 69, 70, 70, 67},
		{87, 92, 94, 96, 99, 99, 99},
		{43, 50, 52, 52, 56},
		{12, 17, 17, 20, 22, 28},
		{67, 73, 74, 77, 81, 82, 84, 86},
		{37, 43, 45, 49, 52, 51},
		{71, 77, 81, 83, 85, 88, 89, 89},
		{5, 11, 15, 16, 19, 23},
		{36, 43, 45, 49, 52, 57},
		{50, 56, 57, 64, 67},
		{49, 54, 61, 64, 61},
		{41, 48, 55, 58, 61, 61},
		{72, 79, 80, 81, 86, 90},
		{57, 62, 64, 65, 68, 70, 77, 83},
		{88, 86, 83, 80, 79, 77, 78},
		{97, 94, 92, 89, 87, 84, 82, 82},
		{74, 73, 71, 68, 65, 63, 62, 58},
		{88, 86, 83, 80, 75},
		{12, 10, 13, 11, 10, 9},
		{33, 32, 35, 33, 30, 29, 27, 28},
		{14, 11, 8, 7, 10, 10},
		{71, 70, 71, 70, 68, 67, 66, 62},
		{20, 18, 16, 18, 15, 13, 12, 7},
		{23, 20, 19, 19, 17},
		{60, 57, 54, 52, 50, 50, 48, 50},
		{72, 71, 71, 69, 69},
		{92, 91, 88, 86, 85, 84, 84, 80},
		{36, 35, 32, 29, 29, 28, 22},
		{41, 38, 34, 31, 30},
		{24, 21, 18, 14, 12, 14},
		{78, 77, 74, 73, 69, 66, 66},
		{30, 28, 27, 24, 20, 17, 13},
		{78, 76, 73, 70, 66, 64, 57},
		{66, 63, 61, 58, 56, 51, 48, 47},
		{19, 17, 15, 12, 11, 5, 7},
		{34, 31, 30, 29, 22, 22},
		{29, 28, 27, 26, 21, 18, 14},
		{94, 93, 92, 85, 79},
		{84, 86, 85, 82, 81, 79, 76, 75},
		{75, 77, 74, 71, 72},
		{72, 75, 74, 73, 72, 72},
		{61, 64, 62, 59, 55},
		{19, 21, 19, 17, 14, 13, 10, 5},
		{4, 5, 4, 7, 4, 3},
		{81, 82, 85, 83, 86},
		{32, 33, 34, 31, 29, 28, 25, 25},
		{57, 58, 55, 52, 55, 51},
		{18, 19, 18, 21, 19, 12},
		{42, 44, 41, 41, 39, 37},
		{82, 84, 82, 81, 81, 79, 80},
		{58, 59, 58, 57, 57, 54, 54},
		{25, 28, 27, 27, 26, 25, 22, 18},
		{97, 99, 96, 93, 90, 87, 87, 82},
		{95, 97, 93, 91, 89, 87},
		{42, 44, 41, 37, 36, 35, 33, 36},
		{35, 37, 33, 31, 28, 28},
		{40, 41, 39, 37, 36, 32, 31, 27},
		{88, 91, 88, 84, 81, 78, 71},
		{93, 94, 93, 92, 86, 85},
		{85, 88, 82, 81, 80, 82},
		{30, 33, 31, 28, 22, 22},
		{58, 59, 54, 53, 52, 50, 48, 44},
		{71, 72, 66, 63, 57},
		{18, 18, 17, 16, 15, 13, 11, 9},
		{47, 47, 44, 43, 40, 38, 36, 39},
		{36, 36, 35, 32, 29, 28, 25, 25},
		{64, 64, 62, 59, 56, 52},
		{99, 99, 96, 94, 92, 90, 84},
		{55, 55, 53, 52, 51, 48, 49, 46},
		{32, 32, 29, 32, 34},
		{94, 94, 91, 89, 86, 83, 85, 85},
		{61, 61, 59, 60, 56},
		{49, 49, 48, 50, 43},
		{8, 8, 5, 5, 2},
		{47, 47, 47, 45, 48},
		{27, 27, 24, 22, 20, 20, 19, 19},
		{99, 99, 96, 93, 91, 91, 90, 86},
		{24, 24, 23, 20, 20, 18, 16, 10},
		{15, 15, 11, 8, 6, 4, 3},
		{72, 72, 68, 65, 64, 63, 65},
		{90, 90, 88, 84, 83, 80, 80},
		{85, 85, 83, 82, 80, 76, 72},
		{66, 66, 62, 59, 52},
		{74, 74, 71, 68, 63, 62, 60, 57},
		{72, 72, 71, 65, 63, 60, 57, 58},
		{68, 68, 67, 64, 61, 54, 54},
		{94, 94, 91, 86, 84, 82, 78},
		{57, 57, 56, 53, 48, 43},
		{94, 90, 89, 86, 83},
		{97, 93, 90, 88, 87, 85, 87},
		{35, 31, 30, 29, 27, 27},
		{72, 68, 66, 64, 60},
		{71, 67, 65, 64, 61, 54},
		{86, 82, 85, 84, 81, 78},
		{35, 31, 32, 30, 32},
		{28, 24, 23, 22, 23, 21, 21},
		{13, 9, 8, 7, 9, 6, 2},
		{81, 77, 74, 72, 75, 69},
		{75, 71, 69, 69, 66, 65, 62, 59},
		{28, 24, 22, 22, 20, 19, 17, 20},
		{37, 33, 30, 29, 27, 27, 27},
		{43, 39, 38, 38, 35, 31},
		{47, 43, 43, 40, 35},
		{22, 18, 15, 11, 8},
		{45, 41, 40, 36, 39},
		{68, 64, 63, 59, 57, 56, 55, 55},
		{36, 32, 29, 27, 23, 20, 16},
		{64, 60, 57, 53, 51, 49, 46, 40},
		{90, 86, 85, 82, 80, 77, 71, 69},
		{25, 21, 18, 15, 13, 8, 10},
		{33, 29, 26, 25, 24, 21, 16, 16},
		{89, 85, 84, 83, 82, 77, 74, 70},
		{83, 79, 76, 75, 72, 71, 66, 59},
		{51, 44, 43, 40, 38, 35, 32},
		{37, 30, 28, 26, 23, 20, 23},
		{79, 72, 71, 68, 66, 63, 61, 61},
		{51, 46, 45, 44, 41, 39, 36, 32},
		{80, 74, 71, 68, 65, 62, 60, 53},
		{92, 85, 88, 85, 84, 82},
		{34, 27, 25, 27, 26, 24, 27},
		{69, 64, 62, 64, 62, 59, 59},
		{31, 24, 21, 20, 18, 19, 15},
		{79, 73, 70, 72, 70, 63},
		{36, 31, 31, 28, 26, 24},
		{23, 18, 17, 17, 15, 18},
		{27, 20, 18, 15, 15, 15},
		{30, 24, 24, 22, 18},
		{56, 49, 46, 46, 44, 39},
		{28, 21, 20, 19, 15, 14, 13},
		{38, 33, 30, 26, 23, 20, 18, 21},
		{54, 48, 47, 44, 42, 38, 38},
		{50, 43, 39, 38, 37, 35, 34, 30},
		{93, 86, 84, 82, 78, 76, 70},
		{60, 54, 47, 46, 44},
		{49, 42, 40, 37, 36, 31, 30, 32},
		{23, 17, 15, 14, 8, 5, 5},
		{97, 92, 91, 90, 84, 83, 79},
		{81, 75, 74, 71, 65, 58},
		{24, 27, 22, 21, 19, 17, 15, 14},
		{62, 59, 57, 56, 53, 53},
		{88, 84, 83, 80, 78, 75, 73, 73},
		{33, 36, 36, 33, 30, 26},
		{66, 63, 66, 68, 69, 67, 70, 73},
		{70, 74, 75, 75, 78, 81, 81},
		{76, 72, 69, 72, 70, 69, 70},
		{30, 30, 28, 27, 23, 21, 16},
		{80, 78, 78, 76, 74, 76},
		{8, 5, 6, 8, 14, 16},
		{86, 87, 90, 92, 96, 97, 94},
		{26, 24, 22, 18, 20},
		{23, 17, 17, 16, 13, 13},
		{78, 74, 74, 71, 68, 65, 61},
		{51, 52, 52, 53, 54, 60},
		{75, 70, 69, 71, 71},
		{32, 36, 36, 37, 39, 43},
		{73, 71, 74, 76, 73, 76, 78, 76},
		{79, 79, 78, 76, 78, 75, 73, 73},
		{68, 67, 68, 66, 69, 70, 70},
		{57, 63, 66, 71, 71},
		{83, 76, 74, 72, 69, 64, 62, 58},
		{45, 50, 48, 50, 51, 48},
		{53, 50, 52, 56, 58, 61, 59},
		{31, 30, 28, 28, 24},
		{91, 87, 84, 81, 77, 76, 76},
		{91, 87, 86, 83, 82, 84, 81, 74},
		{32, 33, 31, 28, 27, 25, 19, 15},
		{25, 31, 33, 35, 37, 40},
		{16, 19, 21, 24, 25, 27, 30, 34},
		{41, 37, 30, 27, 27},
		{57, 56, 63, 64, 64},
		{85, 88, 91, 88, 86, 85, 78},
		{72, 76, 76, 79, 82, 83, 86, 89},
		{4, 7, 9, 12, 12, 15, 16},
		{58, 58, 60, 62, 63, 65, 67, 67},
		{48, 55, 58, 61, 65},
		{76, 76, 79, 77, 78, 78},
		{50, 50, 53, 54, 55, 61, 68},
		{53, 54, 53, 54, 54},
		{73, 71, 71, 73, 77},
		{41, 43, 43, 42, 36},
		{72, 72, 69, 69, 68, 67, 65, 65},
		{25, 22, 20, 20, 18, 16, 14, 8},
		{24, 17, 14, 13, 12, 7, 7},
		{91, 90, 91, 92, 92, 93},
		{5, 8, 10, 11, 16, 17, 20, 20},
		{7, 9, 13, 16, 17, 18, 23},
		{6, 4, 6, 7, 8, 7, 11},
		{73, 66, 64, 63, 62, 64},
		{70, 68, 67, 64, 61, 60, 63},
		{17, 16, 17, 20, 24},
		{69, 73, 75, 76, 80},
		{42, 38, 37, 35, 37, 34, 30},
		{79, 79, 78, 75, 78, 75, 70},
		{36, 35, 40, 43, 47},
		{45, 47, 45, 43, 39, 36, 36},
		{61, 65, 66, 68, 69, 71, 68},
		{84, 77, 77, 75, 73, 72},
		{40, 39, 33, 31, 27},
		{76, 72, 70, 67, 63},
		{23, 19, 16, 15, 14, 9, 6, 9},
		{16, 17, 16, 17, 21},
		{68, 68, 64, 63, 61, 64},
		{77, 77, 79, 82, 83, 86, 87, 84},
		{41, 34, 31, 27, 21},
		{90, 89, 87, 85, 84, 83},
		{77, 75, 73, 72, 71, 70, 69, 66},
		{89, 91, 93, 96, 97, 99},
		{55, 52, 50, 49, 48, 45, 43, 40},
		{75, 74, 71, 69, 67, 65},
		{74, 77, 78, 79, 82, 85, 87, 89},
		{75, 74, 72, 70, 67, 66, 63},
		{32, 29, 26, 23, 20, 19, 16},
		{55, 56, 57, 59, 61, 62, 64},
		{68, 70, 73, 74, 75, 76, 79},
		{30, 29, 26, 23, 22, 21, 19},
		{52, 51, 48, 47, 45},
		{50, 52, 55, 57, 58, 60, 63},
		{59, 60, 63, 64, 66, 68},
		{83, 82, 79, 76, 74, 72, 70},
		{23, 21, 18, 17, 14, 12},
		{32, 33, 35, 38, 39, 42, 45},
		{89, 88, 86, 84, 81, 78, 77, 76},
		{51, 53, 54, 57, 60, 61, 62, 64},
		{19, 18, 15, 12, 10, 8, 6},
		{45, 44, 41, 38, 37, 35, 34},
		{60, 57, 55, 54, 53, 52, 51},
		{55, 57, 58, 60, 63, 66, 67, 69},
		{34, 31, 28, 25, 24, 21, 20},
		{32, 29, 27, 24, 22, 21, 18},
		{30, 28, 25, 22, 19, 16},
		{72, 69, 68, 67, 64, 62},
		{37, 39, 42, 45, 48},
		{62, 65, 67, 70, 73},
		{68, 71, 73, 76, 79, 80, 83, 84},
		{73, 70, 68, 65, 63, 62, 59, 56},
		{58, 55, 53, 51, 49},
		{74, 75, 76, 79, 82, 85},
		{83, 84, 85, 86, 87, 89, 90},
		{88, 91, 94, 95, 98},
		{10, 13, 15, 18, 19},
		{84, 82, 79, 78, 77, 75, 74, 71},
		{27, 30, 33, 35, 37, 40},
		{59, 56, 55, 52, 49},
		{63, 62, 61, 58, 55, 52, 49},
		{93, 90, 89, 87, 86, 85},
		{1, 2, 5, 8, 11},
		{51, 49, 46, 44, 42, 41, 39},
		{64, 65, 66, 69, 72, 75, 76},
		{92, 91, 89, 88, 85, 83, 80},
		{78, 79, 80, 81, 84},
		{40, 39, 36, 34, 32, 31, 30},
		{95, 92, 91, 89, 88, 87, 86, 85},
		{66, 67, 70, 72, 74, 75, 76, 78},
		{45, 48, 49, 51, 54},
		{11, 10, 7, 5, 2},
		{80, 79, 76, 75, 73, 70, 68},
		{33, 31, 28, 26, 24, 21, 18},
		{3, 4, 7, 10, 13},
		{83, 85, 87, 90, 93, 94, 95},
		{31, 32, 33, 35, 37, 38, 41, 42},
		{50, 49, 48, 47, 45, 43, 41},
		{28, 29, 32, 34, 37, 39},
		{77, 79, 81, 83, 85, 86, 89, 91},
		{69, 71, 73, 74, 75, 76, 78, 81},
		{41, 38, 37, 36, 34, 32, 30, 28},
		{86, 85, 84, 81, 80, 79, 77},
		{60, 57, 56, 53, 51},
		{65, 67, 69, 71, 74, 75},
		{85, 86, 88, 91, 93, 95, 97},
		{24, 21, 19, 18, 16, 13},
		{17, 18, 21, 22, 24, 27, 30, 32},
		{77, 79, 80, 81, 82, 83},
		{19, 20, 22, 24, 26, 27, 30},
		{93, 92, 90, 87, 86, 84, 82},
		{84, 86, 88, 91, 94},
		{54, 55, 58, 61, 62},
		{70, 67, 64, 62, 60, 57},
		{27, 30, 33, 34, 35, 37, 40},
		{84, 82, 80, 79, 78},
		{97, 95, 92, 90, 88, 85, 82},
		{48, 47, 46, 44, 42, 39, 38},
		{72, 69, 68, 65, 63, 61},
		{67, 65, 64, 63, 60, 59},
		{85, 83, 82, 80, 79},
		{62, 61, 58, 55, 52},
		{29, 30, 33, 36, 38, 41},
		{74, 75, 78, 79, 81, 82, 83, 86},
		{52, 54, 57, 60, 63},
		{32, 34, 37, 38, 41, 42, 45, 46},
		{14, 15, 18, 20, 22, 25},
		{86, 88, 91, 94, 96, 98, 99},
		{31, 32, 35, 37, 40, 42, 43, 44},
		{3, 4, 7, 10, 13, 15, 16, 17},
		{37, 35, 32, 31, 29, 28, 26, 23},
		{56, 54, 52, 50, 49, 48, 47, 45},
		{57, 60, 62, 64, 66, 68, 69, 72},
		{16, 13, 10, 7, 5},
		{85, 88, 89, 91, 94, 95},
		{42, 40, 37, 36, 33, 31, 28},
		{33, 34, 36, 39, 42, 43},
		{60, 57, 55, 53, 52, 51, 50},
		{42, 44, 46, 47, 49},
		{19, 16, 15, 14, 12, 9},
		{92, 94, 96, 98, 99},
		{73, 71, 68, 66, 64, 63, 60},
		{34, 33, 32, 31, 28, 25},
		{91, 89, 87, 86, 83, 81, 78},
		{7, 8, 11, 13, 16},
		{26, 25, 22, 21, 20, 19, 16},
		{32, 34, 36, 37, 40, 42, 45},
		{73, 70, 67, 66, 65, 63},
		{84, 83, 81, 79, 76, 74},
		{48, 49, 51, 52, 53, 55},
		{53, 55, 56, 58, 59, 60, 62},
		{52, 54, 56, 59, 60, 62},
		{66, 65, 62, 61, 58, 55, 52},
		{83, 86, 87, 88, 91},
		{49, 52, 54, 55, 57, 60, 62, 63},
		{46, 48, 49, 50, 51, 53, 55},
		{40, 41, 42, 43, 45, 48},
		{51, 54, 57, 59, 61, 62},
		{84, 85, 86, 87, 89, 92, 93},
		{8, 6, 5, 3, 1},
		{81, 78, 76, 73, 70},
		{69, 68, 65, 62, 61, 59, 58, 57},
		{58, 55, 52, 51, 49, 46},
		{61, 63, 65, 66, 68},
		{16, 17, 19, 22, 23},
		{22, 20, 18, 17, 16, 15, 14, 12},
		{27, 24, 22, 19, 17},
		{94, 93, 90, 89, 88, 85, 83, 81},
		{42, 41, 38, 35, 34, 32},
		{39, 36, 33, 32, 30, 28},
		{17, 18, 20, 21, 24, 26, 29},
		{37, 38, 39, 40, 41, 44, 46, 49},
		{61, 63, 66, 68, 71, 73, 74},
		{83, 85, 88, 90, 91, 92, 95},
		{31, 34, 35, 38, 39, 40},
		{17, 14, 11, 8, 6, 5, 3},
		{69, 71, 73, 75, 77, 78},
		{64, 67, 70, 73, 74, 75},
		{15, 16, 18, 20, 22, 24, 26, 29},
		{24, 27, 28, 31, 34, 36},
		{74, 75, 76, 77, 78, 79, 82, 83},
		{29, 30, 32, 35, 38, 41, 42, 45},
		{60, 58, 56, 53, 50, 48},
		{27, 24, 22, 20, 19},
		{47, 45, 42, 40, 38, 37},
		{60, 57, 55, 53, 51},
		{67, 66, 63, 62, 59, 57},
		{48, 49, 50, 52, 53},
		{69, 71, 74, 77, 80, 82, 85, 86},
		{62, 65, 66, 68, 69},
		{82, 85, 88, 91, 92},
		{44, 46, 48, 51, 52, 53},
		{96, 94, 92, 89, 86, 83},
		{44, 46, 49, 51, 52},
		{76, 77, 79, 80, 81},
		{78, 75, 74, 73, 72, 71, 68, 65},
		{32, 33, 34, 35, 36, 37},
		{49, 51, 54, 56, 59, 61, 62},
		{63, 65, 67, 69, 70, 72, 73, 76},
		{65, 62, 59, 56, 55},
		{68, 66, 65, 64, 62, 60},
		{67, 65, 63, 61, 59, 56},
		{85, 84, 81, 78, 75, 74, 72, 70},
		{84, 81, 80, 79, 76},
		{55, 54, 52, 51, 50, 47, 45},
		{33, 31, 28, 27, 25, 24, 22},
		{16, 19, 22, 25, 28},
		{76, 77, 80, 81, 84, 86},
		{58, 55, 53, 50, 47},
		{97, 95, 92, 89, 87, 84, 82, 79},
		{42, 43, 45, 47, 50, 51, 54},
		{27, 28, 29, 30, 33, 35, 38, 41},
		{21, 20, 17, 15, 14, 13, 12},
		{32, 34, 37, 38, 41, 44, 46},
		{31, 32, 33, 34, 35, 36, 39, 40},
		{36, 39, 42, 45, 47},
		{62, 65, 67, 69, 71, 72, 74, 75},
		{98, 97, 96, 93, 91},
		{64, 63, 60, 57, 55, 52, 50, 49},
		{53, 54, 55, 58, 61, 63, 66},
		{92, 91, 90, 88, 85},
		{71, 73, 75, 76, 77},
		{79, 82, 85, 87, 89, 92},
		{64, 62, 59, 57, 54, 53, 51},
		{80, 82, 84, 86, 87, 90},
		{22, 24, 27, 28, 30, 33},
		{55, 56, 57, 60, 62, 64, 67},
		{59, 61, 64, 66, 67, 70, 73},
		{8, 10, 13, 15, 17, 19, 22, 25},
		{41, 44, 47, 48, 51},
		{25, 27, 29, 30, 32, 33, 34, 36},
		{63, 60, 58, 56, 55, 53, 51, 50},
		{45, 48, 50, 53, 54, 56},
		{3, 5, 8, 11, 13, 15, 18},
		{30, 33, 34, 35, 37},
		{73, 71, 69, 68, 67, 64, 62, 60},
		{89, 88, 87, 84, 82},
		{45, 43, 40, 38, 36, 34, 33},
		{52, 54, 55, 56, 58, 60, 62, 63},
		{56, 59, 61, 64, 65, 68, 70},
		{53, 52, 51, 49, 48, 47},
		{55, 57, 59, 60, 63, 66},
		{90, 91, 93, 95, 97},
		{62, 65, 66, 68, 70, 71, 74, 77},
		{17, 15, 13, 10, 7, 5, 3},
		{57, 56, 55, 54, 52, 49},
		{60, 57, 56, 54, 53, 52, 50, 47},
		{71, 73, 75, 77, 78, 81},
		{49, 47, 46, 45, 44, 41, 40},
		{47, 46, 45, 44, 43, 40, 38},
		{7, 9, 11, 13, 16, 17, 20},
		{12, 10, 9, 6, 3},
		{21, 19, 16, 14, 13, 12, 11},
		{23, 26, 28, 30, 31},
		{81, 79, 76, 74, 73, 71},
		{56, 59, 61, 63, 65, 68, 71, 74},
		{2, 3, 4, 7, 8, 11},
		{38, 40, 43, 45, 48, 50},
		{26, 24, 22, 21, 20, 17},
		{71, 73, 74, 75, 78},
		{78, 76, 74, 71, 68, 65},
		{85, 86, 87, 89, 91, 93, 96},
		{59, 62, 64, 67, 69},
		{16, 17, 20, 22, 25},
		{91, 88, 86, 83, 82, 79, 76},
		{44, 47, 49, 52, 53, 55, 58, 59},
		{12, 11, 9, 8, 5, 3, 2},
		{86, 87, 90, 92, 95},
		{80, 81, 82, 83, 86, 87, 89},
		{28, 30, 33, 35, 38, 40, 43},
		{90, 88, 86, 84, 82, 79, 78, 75},
		{95, 93, 91, 90, 87},
		{15, 17, 19, 20, 21, 22, 23, 24},
		{72, 71, 70, 68, 67},
		{4, 6, 7, 8, 9, 10, 11, 12},
		{58, 60, 63, 64, 65, 66},
		{80, 83, 86, 88, 89, 90, 91},
		{53, 55, 56, 57, 58, 61, 62},
		{44, 41, 39, 37, 35},
		{71, 73, 74, 76, 77, 78},
		{18, 16, 14, 12, 10, 8},
		{5, 8, 10, 13, 15},
		{74, 77, 79, 81, 83},
		{35, 33, 30, 28, 26},
		{60, 58, 57, 55, 53, 52, 50, 48},
		{12, 10, 7, 6, 3, 1},
		{76, 75, 73, 70, 69, 68},
		{50, 53, 54, 56, 58, 60, 63},
		{38, 36, 33, 31, 28, 27, 26, 23},
		{48, 51, 53, 56, 59, 61},
		{75, 72, 71, 68, 67, 64},
		{27, 25, 22, 20, 17, 16},
		{42, 44, 45, 46, 48, 49},
		{89, 87, 86, 83, 82, 80, 77, 76},
		{97, 96, 93, 91, 89, 86, 84, 83},
		{63, 61, 58, 56, 53, 52},
		{28, 27, 25, 22, 19},
		{82, 80, 78, 77, 74},
		{96, 93, 90, 87, 84, 82, 79},
		{46, 49, 51, 53, 56, 58, 61},
		{69, 67, 65, 64, 61},
		{59, 58, 56, 54, 51, 50, 49, 47},
		{98, 95, 94, 92, 89, 87},
		{59, 58, 57, 54, 52, 50},
		{71, 72, 74, 77, 78, 81},
		{19, 18, 17, 14, 12, 11},
		{81, 80, 79, 76, 74, 72, 71, 69},
		{44, 46, 47, 48, 50, 51},
		{1, 2, 5, 8, 9, 12, 13},
		{77, 79, 82, 85, 88, 89, 92, 95},
		{69, 71, 73, 76, 77},
		{75, 77, 79, 81, 84, 85, 86, 89},
		{4, 6, 8, 9, 11, 12, 15, 18},
		{65, 63, 62, 61, 60, 59, 58},
		{79, 77, 76, 73, 71},
		{59, 57, 56, 55, 53, 51, 49},
		{45, 47, 50, 51, 54},
		{28, 25, 23, 20, 19, 18, 17},
		{46, 43, 41, 38, 35},
		{65, 66, 68, 70, 71},
		{81, 80, 79, 77, 75, 73},
		{82, 84, 85, 86, 89, 90, 93, 95},
		{52, 53, 56, 58, 59},
		{26, 25, 22, 20, 19, 17},
		{81, 82, 84, 86, 87, 88, 91},
		{61, 60, 57, 55, 53},
		{56, 58, 61, 63, 65},
		{74, 72, 70, 68, 67},
		{35, 34, 31, 28, 27, 24, 23},
		{65, 63, 61, 60, 59},
		{88, 89, 90, 92, 94},
		{19, 18, 17, 14, 11, 10, 9, 6},
		{59, 60, 62, 63, 65, 66, 68},
		{66, 64, 63, 62, 60, 59},
		{39, 40, 42, 43, 44},
		{52, 49, 48, 47, 46, 45, 43, 40},
		{40, 42, 44, 45, 46, 49, 50, 53},
		{49, 51, 53, 56, 58},
		{68, 65, 64, 62, 61, 59},
		{72, 73, 76, 77, 80},
		{31, 30, 28, 26, 23, 21, 19},
		{8, 9, 11, 14, 15, 18, 20, 23},
		{2, 5, 7, 8, 11},
		{61, 62, 64, 65, 68, 71, 73, 74},
		{51, 54, 56, 57, 58, 59, 60, 63},
		{19, 16, 14, 12, 10, 9},
		{85, 82, 80, 78, 77, 74, 73, 70},
		{88, 89, 90, 91, 92, 93},
		{37, 38, 39, 42, 45, 46, 48, 50},
		{26, 27, 28, 30, 31, 34},
		{32, 35, 38, 40, 41, 43},
		{37, 38, 40, 41, 44},
		{47, 48, 49, 51, 52, 53, 54},
		{26, 23, 22, 20, 19},
		{99, 97, 96, 93, 92, 89, 86, 83},
		{22, 20, 17, 14, 13},
		{9, 12, 13, 14, 16},
		{26, 24, 22, 21, 20},
		{75, 78, 81, 82, 83, 84, 87, 89},
		{29, 28, 25, 24, 23, 22},
		{98, 95, 92, 89, 88, 85, 82, 81},
		{54, 56, 58, 59, 60, 61},
		{67, 68, 70, 72, 74, 75, 76, 78},
		{48, 46, 45, 42, 41, 40, 37},
		{27, 29, 32, 34, 35},
		{38, 40, 42, 43, 44, 47, 49, 50},
		{59, 60, 63, 64, 66},
		{76, 77, 78, 79, 82, 84, 85, 88},
		{47, 48, 50, 51, 53, 55, 57, 60},
		{5, 6, 7, 10, 13, 15, 17},
		{83, 81, 80, 79, 76, 74, 73, 71},
		{3, 5, 8, 10, 13, 14, 17, 18},
		{37, 36, 33, 32, 31, 30},
		{70, 67, 66, 64, 62},
		{43, 42, 40, 38, 37, 35, 34, 32},
		{18, 20, 21, 24, 25, 28, 30, 33},
		{99, 97, 94, 91, 89, 88, 86, 85},
		{64, 66, 68, 70, 73, 74, 77},
		{24, 27, 28, 29, 31, 32},
		{22, 25, 27, 30, 33},
		{42, 45, 48, 51, 53, 56, 59, 62},
		{87, 88, 90, 93, 94, 95, 97},
		{42, 41, 39, 36, 33, 30},
		{83, 84, 87, 90, 93},
		{17, 20, 23, 24, 27, 28, 31, 34},
		{15, 17, 19, 22, 25},
		{82, 85, 86, 87, 88, 89, 90},
		{1, 2, 5, 7, 8, 11},
		{39, 38, 36, 34, 33},
		{23, 26, 28, 29, 31, 34, 36},
		{58, 59, 60, 61, 62, 65},
		{19, 20, 21, 22, 24, 26, 28},
		{85, 87, 88, 91, 92, 94, 97},
		{96, 93, 90, 89, 88, 85},
		{84, 87, 89, 90, 92},
		{21, 23, 26, 28, 29, 31},
		{51, 53, 54, 57, 59, 60},
		{49, 46, 43, 41, 39, 36, 33},
		{75, 74, 71, 70, 69, 67},
		{59, 62, 63, 65, 66, 68, 71},
		{50, 49, 48, 46, 43, 40, 38, 37},
		{79, 82, 85, 88, 90, 92},
		{74, 72, 71, 69, 66, 63, 61},
		{38, 39, 40, 42, 43, 46, 48},
		{17, 19, 20, 22, 24, 27, 28, 29},
		{46, 48, 49, 51, 52, 53},
		{61, 58, 56, 55, 54, 52, 49, 47},
		{12, 11, 10, 9, 6},
		{74, 76, 77, 80, 81, 82, 85},
		{24, 21, 20, 18, 15, 13},
		{44, 45, 48, 49, 52, 55, 56, 57},
		{29, 31, 34, 37, 39, 41, 42},
		{95, 94, 92, 89, 87, 86, 83, 81},
		{36, 39, 40, 42, 44, 45},
		{63, 64, 66, 69, 71, 74},
		{75, 76, 78, 81, 82},
		{45, 44, 42, 39, 38, 35, 34, 33},
		{37, 38, 41, 44, 47},
		{86, 89, 92, 93, 95, 98},
		{72, 74, 76, 77, 78, 81},
		{48, 45, 42, 39, 38},
		{11, 12, 14, 17, 19, 21, 23},
		{28, 27, 24, 21, 20, 19},
		{65, 66, 69, 71, 73, 75},
		{48, 45, 42, 39, 36},
		{23, 21, 18, 16, 14, 13, 12, 11},
		{87, 84, 81, 80, 79, 76, 75},
		{30, 32, 34, 36, 37},
		{77, 79, 82, 83, 86, 88, 91, 93},
		{86, 88, 90, 92, 93, 95, 97},
		{7, 9, 10, 11, 13, 14, 15, 16},
		{84, 85, 87, 88, 90, 93, 94},
		{71, 69, 67, 65, 63, 62, 60},
		{44, 42, 39, 37, 36, 33, 30},
		{72, 74, 77, 80, 82, 83, 86},
		{79, 76, 73, 70, 67, 66, 63, 62},
		{79, 78, 75, 73, 72},
		{48, 47, 45, 44, 42, 39},
		{19, 18, 15, 14, 13, 12, 11},
		{67, 64, 62, 59, 57, 54, 52, 50},
		{62, 60, 58, 57, 55, 54},
		{31, 30, 28, 25, 23, 20, 19, 17},
		{65, 64, 62, 61, 60, 57},
		{50, 53, 54, 57, 59, 61, 64},
		{30, 28, 25, 22, 19, 16, 13},
		{14, 15, 18, 19, 20, 22, 24},
		{28, 30, 33, 34, 36, 39, 40, 42},
		{50, 53, 56, 57, 60, 61},
		{37, 35, 33, 31, 28},
		{51, 53, 55, 56, 57, 60, 63, 65},
		{44, 42, 41, 39, 37, 34, 33},
		{31, 29, 27, 24, 21, 20, 19, 18},
		{76, 75, 74, 72, 71, 69, 68, 66},
		{63, 64, 66, 67, 69, 70, 71},
		{81, 78, 77, 75, 73},
		{78, 80, 83, 85, 86},
		{79, 82, 84, 87, 89, 91, 94, 95},
		{60, 62, 64, 67, 69, 71},
		{33, 35, 37, 40, 41, 42, 43, 46},
		{87, 88, 91, 92, 94, 95, 96, 99},
		{1, 4, 5, 6, 7, 10},
		{84, 86, 89, 91, 92},
		{18, 21, 23, 24, 27},
		{40, 39, 38, 35, 32},
		{40, 41, 42, 43, 44, 47, 50},
		{16, 14, 12, 11, 8, 6, 5},
		{41, 43, 46, 47, 49, 52, 53},
		{42, 43, 46, 49, 52, 53, 56, 59},
		{72, 70, 69, 66, 64, 61, 59},
		{21, 19, 16, 15, 12, 10},
		{54, 57, 59, 60, 63, 65, 67, 70},
		{60, 59, 57, 54, 52, 50, 48},
		{40, 37, 36, 34, 32},
		{99, 98, 96, 93, 90, 89, 87},
		{55, 52, 51, 48, 46},
		{62, 59, 58, 55, 52},
		{24, 27, 30, 32, 33, 36, 39, 42},
		{79, 77, 75, 74, 73},
		{23, 21, 18, 16, 13, 12, 10, 8},
		{70, 68, 67, 64, 61, 58, 55},
		{20, 23, 24, 25, 28},
		{42, 39, 37, 35, 33},
		{93, 91, 89, 87, 85, 83, 81, 78},
		{41, 42, 44, 45, 48, 51, 54, 57},
		{77, 79, 82, 85, 87, 89, 92, 94},
		{65, 67, 70, 71, 72},
		{24, 22, 20, 17, 16},
		{66, 67, 70, 72, 74, 75, 76, 79},
		{71, 73, 74, 75, 76, 78, 80},
		{76, 78, 79, 80, 83},
		{85, 84, 81, 78, 76, 75},
		{65, 66, 68, 69, 72, 73, 74},
		{82, 85, 88, 90, 92, 95},
		{71, 69, 68, 67, 65, 62, 60},
		{67, 68, 71, 73, 74, 75},
		{76, 79, 81, 83, 86},
		{79, 76, 73, 70, 68, 66, 64},
		{89, 88, 85, 83, 81, 79, 76, 74},
		{49, 51, 53, 55, 56},
		{27, 29, 31, 34, 35, 38, 41, 44},
		{27, 24, 23, 21, 20, 19, 17},
		{12, 15, 17, 19, 20, 22, 24},
		{67, 68, 70, 72, 74},
		{80, 79, 77, 74, 73, 71, 70},
		{42, 45, 46, 48, 51, 53},
		{29, 32, 33, 36, 39, 42, 43},
		{91, 88, 86, 83, 80, 77},
		{33, 35, 36, 37, 40},
		{44, 45, 47, 50, 52},
		{27, 29, 32, 34, 36, 37, 38},
		{42, 45, 47, 48, 49, 50},
		{28, 31, 34, 36, 39},
		{46, 43, 40, 38, 35, 32},
		{27, 30, 33, 36, 38, 40, 43},
		{65, 62, 60, 58, 55},
		{72, 74, 77, 78, 81, 83},
		{85, 88, 91, 94, 96, 98},
		{51, 48, 45, 44, 42, 39, 36},
		{64, 66, 68, 71, 73, 74, 77},
		{96, 95, 93, 92, 90, 87, 86},
		{39, 42, 43, 46, 47, 49},
		{72, 75, 76, 79, 82, 84, 87},
		{57, 55, 54, 52, 50},
		{40, 43, 45, 46, 47, 49},
		{60, 57, 55, 52, 50, 48, 46},
		{72, 69, 67, 66, 63, 62, 61},
		{67, 68, 69, 70, 73, 75},
		{88, 85, 83, 82, 80, 79, 76, 74},
		{22, 19, 17, 16, 13, 12, 10, 8},
		{36, 38, 41, 42, 45, 47, 50},
		{93, 91, 88, 85, 84, 81, 80, 77},
		{65, 67, 70, 71, 73},
		{31, 34, 37, 38, 40, 42, 43},
		{2, 4, 5, 6, 7, 10, 12},
		{16, 15, 13, 10, 8, 6, 4},
		{76, 77, 80, 83, 86},
		{85, 86, 87, 88, 91, 94},
		{36, 33, 31, 29, 26, 25, 23},
		{50, 48, 47, 45, 42, 39, 37, 36},
		{67, 68, 71, 72, 75, 77},
		{79, 76, 74, 72, 69, 68, 65},
		{54, 52, 49, 48, 46, 44},
		{47, 48, 51, 53, 55, 56},
		{74, 71, 69, 68, 67, 65, 63, 61},
		{24, 22, 20, 17, 14, 13},
		{72, 74, 77, 79, 80, 83, 85},
		{28, 25, 24, 23, 20, 17},
		{22, 24, 26, 29, 31, 34},
		{3, 4, 7, 9, 10, 13, 14},
		{21, 18, 17, 14, 11, 10, 9, 8},
		{58, 56, 55, 52, 51, 49, 47},
		{20, 17, 16, 14, 11, 9, 7, 4},
		{34, 37, 39, 40, 42, 45, 46, 49},
		{52, 55, 58, 60, 61, 63},
		{26, 27, 29, 32, 35, 37, 38, 39},
		{59, 62, 65, 68, 70},
		{97, 95, 94, 91, 90, 89, 86},
		{72, 75, 76, 78, 80, 82, 84},
		{72, 69, 66, 65, 63, 62, 61},
		{8, 10, 12, 14, 16, 18, 21, 23},
		{19, 16, 15, 12, 11, 10, 8, 5},
		{47, 49, 52, 55, 56, 59, 62},
		{63, 65, 66, 69, 71, 72, 74, 75},
		{68, 71, 72, 74, 77, 80},
		{61, 62, 64, 66, 68, 70, 72},
		{14, 13, 10, 8, 7},
		{25, 26, 28, 30, 32},
		{68, 66, 64, 63, 61, 59},
		{82, 83, 86, 89, 92, 95, 97},
		{86, 83, 82, 80, 77, 75, 73, 72},
		{68, 70, 73, 74, 75},
		{23, 25, 28, 29, 31, 33, 35, 36},
		{24, 25, 26, 29, 31, 34, 37, 39},
		{71, 74, 77, 80, 82},
		{83, 85, 87, 89, 91, 93},
		{61, 64, 66, 69, 71, 73, 76, 79},
		{84, 82, 80, 78, 77, 75, 72, 69},
		{25, 26, 29, 30, 32, 34},
		{60, 63, 65, 68, 71, 73, 74, 75},
		{73, 70, 69, 68, 67},
		{52, 55, 58, 60, 62},
		{87, 89, 91, 92, 94, 97},
		{23, 22, 20, 17, 16, 14},
		{58, 55, 52, 50, 47, 46, 44, 42},
		{33, 34, 36, 37, 40},
		{29, 31, 33, 36, 39},
		{88, 87, 86, 84, 81},
		{62, 61, 58, 57, 54, 53, 50},
		{11, 13, 15, 17, 19, 21},
		{14, 17, 19, 22, 25, 27, 28, 30},
		{33, 35, 37, 40, 41, 42, 44},
		{53, 50, 48, 45, 44, 41, 39, 38},
		{17, 15, 13, 10, 8, 7, 4, 2},
		{94, 92, 90, 89, 86, 85, 83},
		{57, 56, 53, 51, 50, 47},
		{77, 76, 74, 71, 70},
		{3, 4, 5, 6, 8},
		{87, 90, 93, 94, 97, 98},
		{50, 51, 53, 54, 57, 59},
		{23, 22, 20, 18, 17, 15, 14, 12},
		{88, 87, 84, 82, 80, 78},
		{55, 52, 50, 48, 46},
		{86, 85, 82, 81, 80, 78},
		{80, 78, 75, 73, 71, 70, 69, 66},
		{85, 82, 80, 79, 77},
		{17, 19, 21, 23, 24, 27},
		{3, 4, 7, 9, 12},
		{47, 50, 53, 54, 55},
		{74, 73, 71, 68, 66, 63},
		{78, 79, 81, 82, 85, 86, 89},
		{35, 37, 40, 42, 45},
		{77, 78, 79, 81, 84, 86, 87},
		{63, 65, 67, 69, 72, 74, 75},
		{85, 86, 89, 92, 94},
		{61, 64, 67, 70, 73, 75, 77, 79},
		{39, 36, 33, 32, 30, 29},
		{50, 51, 53, 55, 58, 59, 62, 64},
		{85, 84, 82, 81, 79, 78, 77},
		{10, 12, 13, 16, 19, 20, 22, 25},
		{84, 87, 89, 92, 95},
		{38, 36, 34, 31, 29, 28, 25, 23},
		{74, 76, 79, 81, 84, 86, 89},
		{46, 47, 48, 49, 51, 54, 55},
		{25, 22, 20, 18, 15},
		{22, 24, 26, 28, 31},
		{54, 53, 50, 47, 44, 41, 38, 35},
		{66, 65, 64, 62, 60, 58},
		{80, 77, 74, 72, 69, 66, 65, 62},
		{19, 22, 24, 25, 27, 30},
		{7, 8, 11, 14, 16, 17, 18, 21},
		{46, 44, 43, 40, 37},
		{55, 54, 52, 49, 46},
		{68, 70, 72, 73, 75, 78},
		{33, 30, 27, 26, 23, 21, 20, 19},
		{58, 61, 62, 65, 68, 70, 73, 75},
		{43, 41, 40, 37, 36, 33, 32, 29},
		{33, 32, 30, 29, 27, 26, 23},
		{25, 22, 19, 17, 16},
		{84, 85, 87, 88, 89, 90},
		{34, 37, 40, 42, 44},
		{67, 66, 63, 61, 58, 55, 53, 52},
		{65, 64, 63, 62, 59, 57, 55, 52},
		{74, 75, 76, 77, 80, 81, 83, 84},
		{33, 34, 36, 37, 40, 43, 46, 48},
		{66, 67, 68, 69, 72, 73, 75},
		{86, 89, 92, 93, 95, 97},
		{45, 47, 50, 53, 55, 57, 60},
		{66, 67, 68, 70, 71, 73},
		{37, 39, 40, 41, 43, 44, 46},
		{65, 68, 71, 73, 76},
		{77, 80, 82, 83, 84, 87, 88},
		{79, 77, 74, 73, 70, 69, 66, 64},
		{29, 28, 25, 23, 22, 21, 18},
		{36, 38, 40, 42, 43, 45, 48},
		{26, 23, 21, 18, 17},
		{90, 92, 94, 97, 99},
		{62, 64, 66, 68, 70, 71, 74, 77},
		{85, 83, 81, 80, 77, 76},
		{65, 68, 70, 71, 73, 74, 75},
		{85, 83, 80, 78, 77, 76},
		{68, 67, 65, 62, 59, 57},
		{38, 40, 43, 45, 46, 47, 49},
		{48, 50, 52, 54, 55, 56, 59, 61},
		{64, 63, 62, 59, 58, 56, 54},
		{25, 27, 29, 32, 35, 38},
		{37, 40, 41, 42, 43, 45, 48, 49},
		{64, 63, 61, 58, 56, 53, 50},
		{26, 25, 22, 21, 18, 15, 13, 10},
		{53, 54, 55, 57, 60},
		{42, 39, 36, 34, 31, 30},
		{24, 25, 27, 30, 32},
		{80, 79, 78, 76, 73, 70, 68},
		{18, 20, 23, 26, 27, 29},
		{23, 20, 19, 18, 15},
		{72, 74, 76, 77, 78, 79, 81, 82},
		{30, 28, 27, 25, 22},
		{66, 65, 62, 59, 58, 57, 56},
		{59, 57, 55, 54, 52, 51, 49, 48},
		{74, 75, 76, 77, 79, 82, 85, 86},
		{33, 30, 28, 27, 24, 23, 22},
		{26, 25, 22, 21, 19},
		{70, 71, 74, 76, 77, 79, 82, 85},
		{21, 19, 18, 15, 13, 10, 9},
		{82, 80, 77, 76, 75, 72, 70, 67},
		{90, 88, 86, 84, 81, 79, 77, 74},
		{9, 12, 15, 16, 17},
		{85, 87, 88, 89, 92, 93, 94, 97},
		{63, 60, 57, 54, 53, 52},
		{90, 89, 86, 85, 82, 81, 78},
		{45, 46, 48, 51, 54, 56, 59, 62},
		{96, 93, 90, 89, 88, 86, 83},
		{38, 37, 35, 32, 31, 30},
		{83, 86, 87, 88, 91, 93},
		{35, 36, 39, 42, 43, 45, 46},
		{32, 35, 38, 39, 42, 44, 46},
		{27, 30, 31, 33, 34, 36},
		{50, 53, 55, 58, 61, 64, 65, 66},
		{98, 95, 93, 92, 91, 90, 88, 85},
		{6, 7, 10, 13, 15, 16, 19, 22},
		{15, 18, 20, 21, 24, 27},
		{87, 85, 84, 83, 81, 80, 78},
		{11, 13, 14, 17, 19, 20, 22},
		{28, 27, 24, 22, 21, 19, 16},
		{89, 88, 85, 82, 81, 80, 79},
		{36, 38, 40, 41, 43},
		{63, 62, 61, 58, 55},
		{23, 26, 29, 30, 33, 35, 36, 39},
		{78, 75, 72, 69, 68, 66},
		{64, 61, 60, 59, 58, 57, 56, 54},
		{33, 31, 29, 26, 25},
		{72, 69, 66, 63, 61},
		{3, 4, 7, 9, 10},
		{11, 14, 17, 20, 23},
		{17, 18, 19, 20, 22, 24, 27, 28},
		{88, 89, 92, 95, 96},
		{67, 64, 61, 60, 58, 56, 55},
		{11, 12, 15, 16, 18, 19, 21, 24},
		{7, 10, 12, 13, 15, 18, 21, 23},
		{25, 26, 27, 30, 31},
		{57, 56, 54, 51, 50, 47, 45},
		{68, 71, 72, 75, 77, 78, 81, 83},
		{53, 56, 57, 60, 63, 66, 68, 70},
		{14, 13, 12, 9, 6, 5, 4, 3},
		{59, 56, 53, 52, 50, 48, 45},
		{51, 48, 45, 42, 40, 39, 36},
		{99, 96, 93, 91, 88, 85, 82, 80},
		{22, 21, 18, 16, 14, 13, 12, 10},
	}
	return totalSafeReports(reports)
}

func totalSafeReports(reports [][]int) int {
	result := 0
	for _, report := range reports {
		isSafe := true
		isIncreasing := true
		for i := 0; i < len(report)-1; i++ {
			this, next := report[i], report[i+1]
			delta := delta(this, next)

			// Check delta
			if delta > 3 || delta < 1 {
				isSafe = false
				break
			}

			// Check all increasing or decreasing
			if i == 0 {
				isIncreasing = next > this
				continue
			}

			followsIncOrDecRule := (isIncreasing && next >= this) || (!isIncreasing && next <= this)
			if !followsIncOrDecRule {
				isSafe = false
				break
			}
		}

		if isSafe {
			result++
		}
	}
	return result
}

func delta(a, b int) int {
	result := a - b
	if result < 0 {
		result *= -1
	}
	return result
}
