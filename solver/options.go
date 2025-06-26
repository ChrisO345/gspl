package solver

// TODO: TimeLimit, MaxIterations for LP solvers, IP Solvers need to be implemented first.
type SolverOption struct {
	Tolerance *float64
	// TimeLimit     *time.Duration // Time limit for the solver, 0 means no limit
	MaxIterations *int // Max number of iterations to perform, <= 0 means no limit, used to prevent possible cycling or stalling

	// IP/MIP specific options, will be ignored for LP formulations
	GapSensitivity       *float64           // Gives the minimum gap between the best integer solution and the best bound, 0-1
	UseCuttingPlanes     *bool              // Whether to use cutting planes in the solver
	BranchingStrategy    *BranchingStrategy // e.g., "first-fractional", "most-infeasible"
	HeuristicStrategy    *HeuristicStrategy // e.g., "random", "largest-infeasibility, "none",
	StrongBranchingDepth *int               // Number of strong branching iterations to perform >= 0
	Threads              *int               // Max number of goroutines to spawn, error on negative values, 0 means use all available cores, 1 is single-threaded

	RandomSeed   *int
	SolverMethod *SolverMethod // Method to use for solving, only simplex for now
}

func DefaultSolverOption() SolverOption {
	return SolverOption{
		Tolerance: ptr(1e-6),
		// TimeLimit:     0,      // No time limit
		MaxIterations: ptr(1000), // No iteration limit

		GapSensitivity:       ptr(0.05), // Default gap sensitivity for MIP
		UseCuttingPlanes:     ptr(false),
		BranchingStrategy:    ptr(FirstFractional), // Default branching strategy
		HeuristicStrategy:    ptr(RandomHeuristic), // Default heuristic strategy
		StrongBranchingDepth: ptr(0),               // No strong branching by default
		Threads:              ptr(0),               // Use all available cores

		RandomSeed:   ptr(42),            // Default random seed
		SolverMethod: ptr(SimplexMethod), // Default solver method
	}
}

func (opts *SolverOption) fillDefaults() {
	defaults := DefaultSolverOption()
	if opts == nil {
		opts = &defaults
		return
	}

	if opts.Tolerance == nil {
		opts.Tolerance = defaults.Tolerance
	}
	if opts.MaxIterations == nil {
		opts.MaxIterations = defaults.MaxIterations
	}
	if opts.GapSensitivity == nil {
		opts.GapSensitivity = defaults.GapSensitivity
	}
	if opts.UseCuttingPlanes == nil {
		opts.UseCuttingPlanes = defaults.UseCuttingPlanes
	}
	if opts.BranchingStrategy == nil {
		opts.BranchingStrategy = defaults.BranchingStrategy
	}
	if opts.HeuristicStrategy == nil {
		opts.HeuristicStrategy = defaults.HeuristicStrategy
	}
	if opts.StrongBranchingDepth == nil {
		opts.StrongBranchingDepth = defaults.StrongBranchingDepth
	}
	if opts.Threads == nil {
		opts.Threads = defaults.Threads
	}
	if opts.RandomSeed == nil {
		opts.RandomSeed = defaults.RandomSeed
	}
	if opts.SolverMethod == nil {
		opts.SolverMethod = defaults.SolverMethod
	}
}

func (opts SolverOption) validateSolverOption() error {
	return nil
}

type BranchingStrategy string

const ( // TODO: Add more branching strategies, requires IP solver implementation
	FirstFractional BranchingStrategy = "first-fractional"
	MostFractional  BranchingStrategy = "most-fractional"
	LeastFractional BranchingStrategy = "least-fractional"
	RandomBranching BranchingStrategy = "random-branching"
)

type HeuristicStrategy string

const (
	RandomHeuristic      HeuristicStrategy = "random"
	LargestInfeasibility HeuristicStrategy = "largest-infeasibility"
)

type SolverMethod string

const (
	SimplexMethod SolverMethod = "simplex"
)
