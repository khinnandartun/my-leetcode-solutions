package main

import (
	"fmt"
	"log"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	var result int
	var compareNum int
	compareNum = int(math.Pow(2, 31))

	for j := 0; j <= len(nums)-1; j++ {
		start := j
		end := len(nums) - 1
		for start < end {
			middle := j + 1
			for middle < end {
				//fmt.Println(nums[start], nums[middle], nums[end])
				if nums[start]+nums[middle]+nums[end] == target {
					return target
				} else if nums[start]+nums[middle]+nums[end] < target {
					if ((nums[start]+nums[middle]+nums[end])*(-1) + target) < compareNum {
						compareNum = (nums[start]+nums[middle]+nums[end])*(-1) + target
						result = nums[start] + nums[middle] + nums[end]
					}
				} else {
					//fmt.Println(nums[start]+nums[middle]+nums[end]-target, compareNum)
					if (nums[start] + nums[middle] + nums[end] - target) < compareNum {
						compareNum = nums[start] + nums[middle] + nums[end] - target
						result = nums[start] + nums[middle] + nums[end]
					}
				}
				//fmt.Println("result  ", result)
				middle = middle + 1
			}
			end = end - 1
		}
	}
	return result
}
func main() {
	//fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, 100))
	// fmt.Println(threeSumClosest([]int{558, 316, -411, 160, 801, 568, -124, -589, 32, 897, -33, -767, -528, -180, 916, 813, 351, 642, -373, -919, 666, 973, -165, 831, -67, -934, -659, -18, 273, 201, 728, 988, -926, 409, -573, 575, -502, 745, 724, 989, -464, 903, 975, 980, 824, -197, -261, -761, 966, 799, -379, 96, 9, -680, -15, 476, 220, -647, 365, 518, -479, -443, 337, -364, 968, -617, 862, -281, -936, -526, 196, 829, -191, 643, -473, 557, -870, 553, -506, 774, 784, -344, -452, 510, 219, -785, -1, 711, -759, -830, 10, 612, -450, -784, 53, 976, -314, -908, 463, -408, -846, 261, 689, -856, -687, -949, -163, -621, -233, 847, -682, -805, -711, 286, 40, -831, -12, 952, -878, -226, 739, 11, -342, 74, -933, -770, -840, 265, 702, 572, -453, -295, 704, -249, -835, 191, 404, 984, -820, 321, 632, -689, 285, -877, -643, -451, -625, 84, 889, 620, -658, 861, -397, -952, 695, -386, 31, 827, -539, -350, 588, 846, -142, 314, 909, 937, 625, -230, -553, 403, -763, 413, 904, -994, 272, -426, 104, -715, -159, -270, 716, 819, 806, 891, -47, -100, 440, -339, 918, -577, 508, -554, -478, 120, -943, 25, -600, -626, 336, -567, 473, 531, 195, -259, -267, -883, 450, 170, 733, 491, 602},
	// 	-8224))
	log.Println("start")
	fmt.Println(threeSumClosest([]int{
		228, 306, -217, -71, 770, -128, 60, 277, 230, -383, -161, 88, 841, 373, 493, -833, 898, 701, -252, 884, -915, -809, -743, 724, 324, -179, -375, 497, -844, 674, 601, -229, -612, 372, -64, -781, -87, 951, 741, -577, 803, 700, 500, -866, -97, -559, 697, -566, -51, 38, 166, -430, 317, 825, 962, 320, -705, -834, -277, -676, 45, 985, 40, 441, 801, 214, -583, 474, -431, -700, 745, 836, -444, -296, 400, 369, 206, 636, 42, -463, 14, -156, -739, -943, 662, -70, -126, 957, 793, 602, -65, 169, -978, -608, 448, 845, -696, -20, 445, 257, 552, -800, 140, -379, 621, 935, -801, -232, 524, 48, -761, -826, -827, -518, -841, -358, -508, -531, 447, 208, -511, -385, -451, 575, -121, 344, 56, -225, -41, 154, 812, -640, 192, 95, -99, 287, -878, 330, 449, 36, -814, -522, -813, 970, 588, 574, 541, -581, 131, 273, -440, 668, -856, -488, 508, -891, -627, 689, -868, -108, 274, 796, -181, -709, 165, -455, 327, 710, -153, -895, 922, 416, -62, -748, 910, 877, 997, -685, 499, -873, 115, -266, 312, -312, -249, -586, -546, 654, -904, 528, 25, -728, 133, 978, -392, -609, -172, 376, 249, 529, -937, -388, 15, 89, -513, 807, -946, 886, 82, 735, 761, 729, 959, 13, 632, -256, -344, 831, -343, 483, -556, -354, 325, 217, -76, -691, -319, 696, -506, 269, -485, -294, 965, -775, 501, -807, 76, 179, -538, 643, -287, -958, 681, 940, -784, 768, -36, 535, -22, 790, 232, 238, 79, 576, -452, -400, 607, 68, 897, -850, -698, 83, -733, 579, 486, -495, 869, -493, 495, -113, -177, -118, -681, -248, 963, 480, -823, -913, -510, 386, 699, -462, 655, 227, -766, 252, -621, 815, -334, -597, -330, -962, -23, 446, 800, -567, -314, -877, -163, 640, 433, -133, -264, -626, -973, 290, 924, -105, -885, -629, -136, -107, 644, 543, -883, -774, -145, -899, 7, 71, 954, -929, 648, 932, -116, 364, -253, 426, 818, -438, 794, 934, -778, -235, 851, 811, 155, -257, -953, -928, 998, 615, 350, 926, 92, 769, -727, 345, 109, -684, -336, 346, 417, 0, -63, 973, -604, -56, 200, 20, 510, 161, 685, -2, 218, -776, 917, 308, -544, 357, -724, 519, -18, 16, -433, -832, 590, -272, 93, 694, -863, -554, -865, 278, 181, 829, -475, 172, -934, -189, 739, 240, 980, -373, -892, -515, -491, -712, -159, 225, -639, 665, -351, 949, -704, 933, -57, 65, 742, -387, -884, 340, -964, -686, -80, 174, 390, 91, 880, 537, 113, -407, -912, -94, 788, -648, 580, 631, 314, -199, 432, -465, 852, 104, -932, -384, -91, -428, 708, 450, -646, 421, 805, -718, -365, 260, -135, -979, 971, -726, 743, 868, 928, 434, 725, 191, -643, 792, 874, 331, 332, 703, -137, -349, 3, -50, -951, 559, -467, 764, 122, -47, -472, -647, -535, 758, 144, 929, -352, -526, 53, 589, 675, -215, -505, -751, -871, -947, -703, -34, 150, -1000, -476, -193, -441, 319, 243, -688, -516, 63, -168, 375, 195, -501, 167, 861, -504, -234, 349, 586, 901, 722, 671, -821, -83, -127, 454, 494, -984, -398, 946, 629, -520, -353, 718, 148, 969, -39, -763, 146, 715, 751, 359, -282, -864, 894, 313, -674, -157, 171, 827, 866, -764, -171, 479, -320, 778, 628, -447, -507, -638, -716, -3, -547, -332, -8, 714, 378, -738, 705, 335, 626, -404, -289, -957, 17, 566, 177, -54, -860, 173, -982, -140, 616, 863, 43, 134, 842, -812, -817, 864, 496, 462, 415, 389, -158, -777, 717, -887, -92, 209, -988, -29, 410, -975, -303, -561, 64, 347, 748, 713, -25, -889, 904, -794, 577, -666, -300, 862, -449, -242, -362, -907, -618, 723, 283, 112, -424, 272, 487, -598, 719, 299, -689, -908, 856, -316, -243, 279, -767, -558, 156, 944, 18, -106, -714, 360, 342, 670, 905, 370, 994, -869, 780, 887, -283, -75, 765, 266, -468, 35, -251, 385, 211, 504, 947, 49, 244, 86, 534, 542, -999, -1, -411, -269, 51, 268, -81, -996, -98, -762, -149, 979, -308, -990, -304, 406, -808, -203, 391, 21, -896, 328, 489, 420, 213, 733, 545, -478, 117, -768, 210, -134, 10, 597, -672, 224, 285, -230, -155, -560, 298, 516, 753, -687, -881, -187, 407, -377, -971, 720, 571, 747, -593, 334, -741, -569, -634, 967, 539, -550, 419, -454, -514, 429, -750, 687, 736, -486, 94, -965, -997, -616, -32, 212, 937, -960, 832, 883, -412, -614, -24, 799, 187, -579, 596, 564, 521, 66, -148, -548, -209, 185, -798, 247, -938, 658, -576, 673, -747, -785, 57, 706, -786, 39, -539, -858, -35, -717, -805, 584, 142, -936, 197, 907, -49, 820, -854, 354, -389, 205, 762, 952, -43, 105, 73, 502, -842, 54, -152, -870, -720, 814, -280, 808, 507, 490, -231, -948, -673, 610, 903, 515, 311, 756, 377, -980, -174, -654, 773, -631, -90, -357, 728, 292, 223, -201, -752, 757, 966, 233, 824, 338, 264, 996, 5, -12, 248, 513, 953, -995, -167, 821, -408, 50, -680, 915, -662, 591, -216, -131, -900, 611, 22, -497, 98, -802, -825, -196, -16, 59, 467, 603, -939, 639, -499, 37, 451, -744, -429, -233, -442, 704, -998, 547, -422, 159, 882, 1000, 553, -525, 304, -653, -170, 288, 678, -448, -521, 460, 810, -255, 459, -123, -810, 463, 565, -737, -555, 289, 199, 392, 860, -89, -124, 296, -40, -78, -950, -390, 902, 184, -397, 895, -578, 514, -903, -532, 444, -512, 374, 931, 511, 286, -587, 637, 614, 634, 41, 606, -420, -461, 28, -779, -607, 690, 481, 653, -58, -443, -859, -879, 190, -905, -848, -394, -790, 925, -855, -402, 968, -265, 813, -773, -839, 81, 182}, 3647))
	log.Println("stop")
}
