import type {
  GraphHealthFilterable,
} from "./graph-health-filters";
import type {
  SemanticFilterable,
} from "./semantic-filters";

export interface SigmaNodeAttributes
  extends GraphHealthFilterable,
    SemanticFilterable {
  x: number;
  y: number;
  size: number;
  type?: string;
  color: string;
  label: string;
  nodeType: string;
  rawNodeType?: string;
  filePath: string;
  startLine?: number;
  endLine?: number;
  hidden?: boolean;
  zIndex?: number;
  highlighted?: boolean;
  mass?: number;
  community?: number;
  communityColor?: string;
  confidence?: string;
  appLayerRing?: string;
  islandKey?: string;
  appLayerRingCenterX?: number;
  appLayerRingCenterY?: number;
}

export interface SigmaEdgeAttributes {
  size: number;
  color: string;
  relationType: string;
  hidden?: boolean;
  zIndex?: number;
}

